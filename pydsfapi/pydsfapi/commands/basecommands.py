"""
This module contains all basic commands to be sent to the server.

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
from enum import IntEnum, Enum
from .codechannel import CodeChannel


class BaseCommand:
    """Base class of a command."""
    @classmethod
    def from_json(cls, data):
        """Deserialize an instance of this class from a JSON deserialized dictionary"""
        return cls(**data)

    def __init__(self, command: str, **kwargs):
        self.command = command
        for key, value in kwargs.items():
            self.__dict__[key] = value


ACKNOWLEDGE = BaseCommand('Acknowledge')
CANCEL = BaseCommand('Cancel')
IGNORE = BaseCommand('Ignore')
GET_MACHINE_MODEL = BaseCommand('GetObjectModel')
GET_OBJECT_MODEL = BaseCommand('GetObjectModel')
SYNC_MACHINE_MODEL = BaseCommand('SyncObjectModel')
SYNC_OBJECT_MODEL = BaseCommand('SyncObjectModel')
LOCK_MACHINE_MODEL = BaseCommand('LockObjectModel')
LOCK_OBJECT_MODEL = BaseCommand('LockObjectModel')
UNLOCK_MACHINE_MODEL = BaseCommand('UnlockObjectModel')
UNLOCK_OBJECT_MODEL = BaseCommand('UnlockObjectModel')


class HttpEndpointType(str, Enum):
    """Enumeration of supported HTTP request types"""
    GET = 'GET'
    POST = 'POST'
    PUT = 'PUT'
    PATCH = 'PATCH'
    TRACE = 'TRACE'
    DELETE = 'DELETE'
    OPTIONS = 'OPTIONS'
    WebSocket = 'WebSocket'


def add_http_endpoint(endpoint_type: HttpEndpointType, namespace: str, path: str,
                      is_upload_request: bool):
    """
    Register a new HTTP endpoint via DuetWebServer.
    This will create a new HTTP endpoint under /machine/{Namespace}/{EndpointPath}.
    Returns a path to the UNIX socket which DuetWebServer will connect to whenever a matching
    HTTP request is received. A plugin using this command has to open a new UNIX socket with
    the given path that DuetWebServer can connect to
    """
    return BaseCommand(
        'AddHttpEndpoint', **{
            'EndpointType': endpoint_type,
            'Namespace': namespace,
            'Path': path,
            'IsUploadRequest': is_upload_request,
        })


def remove_http_endpoint(endpoint_type: HttpEndpointType, namespace: str, path: str):
    """
    Remove an existing HTTP endpoint.
    Returns true if the endpoint could be successfully removed
    """
    return BaseCommand('RemoveHttpEndpoint'**{
        'EndpointType': endpoint_type,
        'Namespace': namespace,
        'Path': path
    })


class AccessLevel(str, Enum):
    """Defines what a user is allowed to do"""
    ReadOnly = 'ReadOnly'
    ReadWrite = 'ReadWrite'


class SessionType(str, Enum):
    """Types of user sessions"""
    Local = 'Local'
    HTTP = 'HTTP'
    Telnet = 'Telnet'


def add_user_session(access: AccessLevel, tpe: SessionType, origin: str, origin_port: int):
    """
    Register a new user session.
    Returns the ID of the new user session
    """
    return BaseCommand(
        'AddUserSession', **{
            'AccessLevel': access,
            'SessionType': tpe,
            'Origin': origin,
            'OriginPort': origin_port
        })


def remove_user_session(session_id: int):
    """Remove an existing user session"""
    return BaseCommand('RemoveUserSession', **{'Id': session_id})


def evaluate_expression(channel: CodeChannel, expression: str):
    """
    Evaluate an arbitrary expression on the given channel in RepRapFirmware.
    Do not use this call to evaluate file-based and network-related fields because the
    DSF and RRF models diverge in this regard.
    """
    return BaseCommand('EvaluateExpression', **{'Channel': channel, 'Expression': expression})


def flush(channel: CodeChannel):
    """Create a Flush command"""
    return BaseCommand('Flush', **{'Channel': channel})


def get_file_info(file_name: str):
    """Create a GetFileInfo command"""
    return BaseCommand('GetFileInfo', **{'FileName': file_name})


def resolve_path(path: str):
    """Create a ResolvePath command"""
    return BaseCommand('ResolvePath', **{'Path': path})


def simple_code(code: str, channel: CodeChannel):
    """Create a simple G/M/T code command"""
    return BaseCommand('SimpleCode', **{'Code': code, 'Channel': channel})


def patch_object_mode(key: str, patch: str):
    """
    Apply a full patch tot he object model. May be used only in non-SPI mode
    """
    return BaseCommand('PatchObjectModel', **{'Key': key, 'Patch': patch})


def set_machine_model(property_path: str, value: str):
    """
    Set an atomic property in the machine model.
    Make sure to acquire the read/write lock first! Returns true if the field could be updated
    """
    return BaseCommand('SetObjectModel', **{'PropertyPath': property_path, 'Value': value})


def set_update_status(updating: bool):
    """
    Override the current status as reported by the object model when
    performing a software update.
    """
    return BaseCommand('SetUpdateStatus', **{'Updating': updating})


def install_plugin(plugin_file: str):
    """
    Install or upgrade a plugin
    """
    return BaseCommand('InstallPlugin', **{'PluginFile': plugin_file})


def set_plugin_data(plugin: str, key: str, value: str):
    """
    Set custom plugin data in the object model.
    May be used to update only the own plygin data unless the plugin has the ManagePlugins permission
    """
    return BaseCommand('SetPluginData', **{'Plugin': plugin, 'Key': key, 'Value': value})


def start_plugin(plugin: str):
    """
    Start a plugin
    """
    return BaseCommand('StartPlugin', **{'Plugin': plugin})


def stop_plugin(plugin: str):
    """
    Stop a plugin
    """
    return BaseCommand('StopPlugin', **{'Plugin': plugin})


def uninstall_plugin(plugin: str):
    """
    Uninstall a plugin
    """
    return BaseCommand('UninstallPlugin', **{'Plugin': plugin})


class MessageType(IntEnum):
    """Type of Resolve message"""
    Success = 0
    Warning = 1
    Error = 2


def resolve_code(rtype: MessageType, content: str):
    """Create a Resolve message"""
    return BaseCommand('Resolve', **{'Type': rtype, 'Content': content})
