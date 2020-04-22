"""
connections provides different classes for connections

    Python interface to DuetSoftwareFramework
    Copyright (C) 2020 Duet3D

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Lesser General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Lesser General Public License for more details.

    You should have received a copy of the GNU Lesser General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
"""
from enum import Enum
import asyncio
import json
import os
import socket
from concurrent.futures import ThreadPoolExecutor
from .commands import responses, basecommands, code, result, codechannel
from .commands.basecommands import MessageType
from .initmessages import serverinitmessage, clientinitmessages
from .model import parsedfileinfo, machinemodel

SOCKET_DIRECTORY = "/var/run/dsf"
SOCKET_FILE = "dcs.sock"
FULL_SOCKET_PATH = SOCKET_DIRECTORY + "/" + SOCKET_FILE
DEFAULT_BACKLOG = 4


class TaskCanceledException(Exception):
    """Exception returned by the server if the task has been cancelled remotely"""


class InternalServerException(Exception):
    """Exception returned by the server for an arbitrary problem"""

    def __init__(self, command, error_type: str, error_message: str):
        super().__init__('Internal Server Exception')
        self.command = command
        self.error_type = error_type
        self.error_message = error_message


class HttpResponseType(str, Enum):
    """Enumeration of supported HTTP responses"""
    StatusCode = 'StatusCode'
    PlainText = 'PlainText'
    JSON = 'JSON'
    File = 'File'


class ReceivedHttpRequest:
    """Notification sent by the webserver when a new HTTP request is received"""
    @classmethod
    def from_json(cls, data):
        """Deserialize an instance of this class from deserialized JSON dictionary"""
        return cls(**data)

    def __init__(self, sessionId: int, queries: dict, headers: dict, contentType: str, body: str):
        self.session_id = sessionId
        self.queries = queries
        self.headers = headers
        self.content_type = contentType
        self.body = body


class HttpEndpointConnection:
    """Connection class for dealing with requests received from a custom HTTP endpoint"""

    def __init__(self, reader, writer, is_websocket: bool, debug: bool = False):
        """Constructor for a new connection dealing with a single HTTP endpoint request"""
        self.reader = reader
        self.writer = writer
        self.is_websocket = is_websocket
        self.debug = debug

    def close(self):
        """Close the connection"""
        self.writer.close()

    async def read_request(self):
        """
        Read information about the last HTTP request.
        Note that a call to this method may fail!
        """
        return await self.receive(ReceivedHttpRequest)

    async def send_response(self,
                            status_code: int = 204,
                            response: str = "",
                            response_type: HttpResponseType = HttpResponseType.StatusCode):
        """
        Send a simple HTTP response to the client and dispose
        this connection unless it is a WebSocket.
        """
        try:
            await self.send({
                'StatusCode': status_code,
                'Response': response,
                'ResponseType': response_type
            })
        finally:
            # Close this connection automatically if only one response can be sent
            if not self.is_websocket:
                self.close()

    async def receive(self, cls):
        """Receive a deserialized object"""
        json_string = await self.receive_json()
        return cls.from_json(json.loads(json_string))

    async def receive_json(self):
        """Receive a JSON object"""
        json_string = (await self.reader.read(32 * 1024)).decode('utf8')
        if self.debug:
            print('recv: {0}'.format(json_string))
        return json_string

    async def send(self, obj):
        """Send an arbitrary object"""
        json_string = json.dumps(obj, default=lambda o: o.__dict__)
        if self.debug:
            print('send: {0}'.format(json_string))
        self.writer.write(json_string.encode('utf8'))
        await self.writer.drain()


class HttpEndpointUnixSocket:
    """Class for dealing with custom HTTP endpoints"""

    def __init__(self,
                 endpoint_type: basecommands.HttpEndpointType,
                 namespace: str,
                 path: str,
                 socket_path: str,
                 backlog: int = DEFAULT_BACKLOG,
                 debug: bool = False):
        """Open a new UNIX socket on the given file path"""
        self.endpoint_type = endpoint_type
        self.namespace = namespace
        self.endpoint_path = path
        self.socket_path = socket_path
        self.backlog = backlog
        self.handler = None
        self.debug = debug
        self._loop = None
        self._server = None

        try:
            os.remove(self.socket_path)
        except:
            # We don't care if the file was missing
            # TODO: should we care about deletion failed?
            pass

        self.executor = ThreadPoolExecutor(max_workers=1)
        self.event_loop = self.executor.submit(self.start_connection_listener)

    def close(self):
        """Close the socket connection"""
        if self._loop is not None:
            # TODO: this enables correctly ending the loop. Why?
            self._loop.set_debug(True)
            self._server.close()
            self._loop.stop()
        self.event_loop.cancel()
        self.executor.shutdown(wait=False)
        try:
            os.remove(self.socket_path)
        except:
            pass

    def set_endpoint_handler(self, handler):
        """Set the handler to handle client connections"""
        self.handler = handler

    def start_connection_listener(self):
        try:
            self._loop = asyncio.new_event_loop()
            self._server = asyncio.start_unix_server(self.handle_connection,
                                                     self.socket_path,
                                                     backlog=self.backlog)
            self._loop.create_task(self._server)
            self._loop.run_forever()
        finally:
            self._loop.close()

    async def handle_connection(self, reader, writer):
        """Handle incoming UNIX socket connections (HTTP/WebSocket requests)"""
        http_endpoint_connection = HttpEndpointConnection(
            reader,
            writer,
            self.endpoint_type == basecommands.HttpEndpointType.WebSocket,
            debug=self.debug)
        if self.handler is not None:
            # Invoke the event handler and forward the wrapped connection for dealing
            # with a single endpoint connection.
            # Note that the event delegate is responsible for disposing the connection!
            await self.handler(http_endpoint_connection)
        else:
            await http_endpoint_connection.send_response(500, "No event handler registered")
            http_endpoint_connection.close()


class BaseConnection:
    """
    Base class for connections that access the control server via the Duet API
    using a UNIX socket
    """

    def __init__(self, debug: bool = False):
        self.debug = debug
        self.socket = None
        self.id = None

    def connect(self, init_message: clientinitmessages.ClientInitMessage, socket_path: str):
        """Establishes a connection to the given UNIX socket file"""

        self.socket = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        self.socket.connect(socket_path)
        self.socket.setblocking(True)
        server_init_message = serverinitmessage.ServerInitMessage.from_json(
            json.loads(self.socket.recv(50).decode('utf8')))
        if not server_init_message.is_compatible():
            raise serverinitmessage.IncompatibleVersionException(
                'Incompatible API version (need {0}, got {1})'.format(
                    server_init_message.PROTOCOL_VERSION, server_init_message.version))
        self.id = server_init_message.id
        self.send(init_message)

        response = self.receive_response()
        if not response.success:
            raise Exception('Could not set connection type {0} ({1}: {2})'.format(
                init_message.mode, response.error_type, response.error_message))

    def close(self):
        """Closes the current connection and disposes it"""
        if self.socket is not None:
            self.socket.close()
            self.socket = None

    def perform_command(self, command, cls=None):
        """Perform an arbitrary command"""
        self.send(command)

        response = self.receive_response()
        if response.success:
            if cls is not None and response.result is not None:
                response.result = cls.from_json(response.result)
            return response

        if response.error_type == 'TaskCanceledException':
            raise TaskCanceledException(response.error_message)

        raise InternalServerException(
            command, response.error_type, response.error_message)

    def send(self, msg):
        """Serialize an arbitrary object into JSON and send it to the server plus NL"""
        json_string = json.dumps(msg, default=lambda o: o.__dict__)
        if self.debug:
            print('send: {0}'.format(json_string))
        self.socket.sendall(json_string.encode('utf8'))

    def receive(self, cls):
        """Receive a deserialized object from the server"""
        json_string = self.receive_json()
        return cls.from_json(json.loads(json_string))

    def receive_response(self):
        """Receive a base response from the server"""
        json_string = self.receive_json()
        return json.loads(json_string, object_hook=responses.decode_response)

    def receive_json(self):
        """Receive the JSON response from the server"""
        json_string = self.socket.recv(32 * 1024).decode('utf8')
        if self.debug:
            print('recv: {0}'.format(json_string))
        return json_string


class BaseCommandConnection(BaseConnection):
    """Base connection class for sending commands to the control server"""

    def flush(self, channel: codechannel.CodeChannel = codechannel.CodeChannel.SBC):
        """Wait for all pending codes of the given channel to finish"""
        return self.perform_command(basecommands.flush(channel))

    def add_http_endpoint(self,
                          endpoint_type: basecommands.HttpEndpointType,
                          namespace: str,
                          path: str,
                          backlog: int = DEFAULT_BACKLOG):
        """Add a new third-party HTTP endpoint in the format /machine/{ns}/{path}"""
        res = self.perform_command(
            basecommands.add_http_endpoint(endpoint_type, namespace, path))
        socket_path = res.result
        return HttpEndpointUnixSocket(endpoint_type, namespace, path, socket_path, backlog,
                                      self.debug)

    def add_user_session(self,
                         access: basecommands.AccessLevel,
                         tpe: basecommands.SessionType,
                         origin: str,
                         origin_port: int = None):
        """Add a new user session"""
        if origin_port is None:
            origin_port = os.getpid()

        res = self.perform_command(basecommands.add_user_session(
            access, tpe, origin, origin_port))
        return int(res.result)

    def get_file_info(self, file_name: str):
        """Parse a G-code file and returns file information about it"""
        res = self.perform_command(basecommands.get_file_info(file_name),
                                   parsedfileinfo.ParsedFileInfo)
        return res.result

    def get_machine_model(self):
        """Retrieve the full object model of the machine."""
        res = self.perform_command(
            basecommands.GET_MACHINE_MODEL, machinemodel.MachineModel)
        return res.result

    def get_serialized_machine_model(self):
        """Optimized method to directly query the machine model UTF-8 JSON"""
        self.send(basecommands.GET_MACHINE_MODEL)
        return self.receive_json()

    def lock_machine_model(self):
        """
        Lock the machine model for read/write access.
        It is MANDATORY to call unlock_machine_model when write access has finished
        """
        return self.perform_command(basecommands.LOCK_MACHINE_MODEL)

    def perform_code(self, cde: code.Code):
        """Execute an arbitrary pre-parsed code"""
        res = self.perform_command(cde, result.CodeResult)
        return res.result

    def perform_simple_code(self, cde: str, channel: codechannel.CodeChannel = codechannel.CodeChannel.DEFAULT_CHANNEL):
        """Execute an arbitrary G/M/T-code in text form and return the result as a string"""
        res = self.perform_command(basecommands.simple_code(cde, channel))
        return res.result

    def remove_http_endpoint(self, endpoint_type: basecommands.HttpEndpointType, namespace: str,
                             path: str):
        """Remove an existing HTTP endpoint"""
        res = self.perform_command(basecommands.remove_http_endpoint(endpoint_type, namespace,
                                                                     path))
        return res.result

    def remove_user_session(self, session_id: int):
        """Remove an existing HTTP endpoint"""
        res = self.perform_command(
            basecommands.remove_user_session(session_id))
        return res.result

    def resolve_path(self, path: str):
        """Resolve a RepRapFirmware-style file path to a real file path"""
        return self.perform_command(basecommands.resolve_path(path))

    def set_machine_model(self, path: str, value: str):
        """
        Set a given property to a certain value.
        Make sure to lock the object model before calling this
        """
        return self.perform_command(basecommands.set_machine_model(path, value))

    def sync_machine_model(self):
        """Wait for the full machine model to be updated from RepRapFirmware"""
        return self.perform_command(basecommands.SYNC_MACHINE_MODEL)

    def unlock_machine_model(self):
        """Unlock the machine model again"""
        return self.perform_command(basecommands.UNLOCK_MACHINE_MODEL)


class CommandConnection(BaseCommandConnection):
    """Connection class for sending commands to the control server"""

    def connect(self, socket_path: str = FULL_SOCKET_PATH):
        """Establishes a connection to the given UNIX socket file"""
        return super().connect(clientinitmessages.command_init_message(), socket_path)


class InterceptConnection(BaseCommandConnection):
    """Connection class for intercepting G/M/T-codes from the control server"""

    def __init__(self, interception_mode: clientinitmessages.InterceptionMode, debug: bool = False):
        super().__init__(debug)
        self.interception_mode = interception_mode

    def connect(self, socket_path: str = FULL_SOCKET_PATH):
        """Establishes a connection to the given UNIX socket file"""
        iim = clientinitmessages.intercept_init_message(self.interception_mode)

        return super().connect(iim, socket_path)

    def receive_code(self):
        """Wait for a code to be intercepted and read it"""
        return self.receive(code.Code)

    def cancel_code(self):
        """Instruct the control server to cancel the last received code (in intercepting mode)"""
        self.send(basecommands.CANCEL)

    def ignore_code(self):
        """Instruct the control server to ignore the last received code (in intercepting mode)"""
        self.send(basecommands.IGNORE)

    def resolve_code(self, rtype: MessageType = MessageType.Success, content: str = None):
        """
        Instruct the control server to resolve the last received code with the given
        message details (in intercepting mode)
        """
        self.send(basecommands.resolve_code(rtype, content))


class SubscribeConnection(BaseConnection):
    """Connection class for subscribing to model updates"""

    def __init__(self,
                 subscription_mode: clientinitmessages.SubscriptionMode,
                 filter_str: str = "",
                 debug: bool = False):
        super().__init__(debug)
        self.subscription_mode = subscription_mode
        self.filter_str = filter_str

    def connect(self, socket_path: str = FULL_SOCKET_PATH):
        """Establishes a connection to the given UNIX socket file"""
        sim = clientinitmessages.subscibe_init_message(
            self.subscription_mode, self.filter_str)

        return super().connect(sim, socket_path)

    def get_machine_model(self):
        """
        Retrieves the full object model of the machine
        In subscription mode this is the first command that has to be called once a
        ConnectionAbortedError has been established.
        """
        machine_model = self.receive(machinemodel.MachineModel)
        self.send(basecommands.ACKNOWLEDGE)
        return machine_model

    def get_serialized_machine_model(self):
        """
        Optimized method to query the machine model UTF-8 JSON in any mode.
        May be used to get machine model patches as well.
        """
        machine_model_json = self.receive_json()
        self.send(basecommands.ACKNOWLEDGE)
        return machine_model_json

    def get_machine_model_patch(self):
        """
        Receive a (partial) machine model update.
        If the subscription mode is set to SubscriptionMode.PATCH new update patches of
        the object model need to be applied manually. This method is intended to receive
        such fragments.
        """
        patch_json = self.receive_json()
        self.send(basecommands.ACKNOWLEDGE)
        return patch_json
