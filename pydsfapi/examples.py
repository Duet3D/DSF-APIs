#!/usr/bin/env python3

from commands import basecommands, code
import sys
import time
import pydsfapi
from initmessages.clientinitmessages import InterceptionMode, SubscriptionMode

def intercept():
    intercept_connection = pydsfapi.InterceptConnection(InterceptionMode.PRE, debug=True)
    intercept_connection.connect()

    try:
        while True:
            cde = intercept_connection.receive_code()
            if cde.type == code.CodeType.MCode:
                print(cde, cde.flags)
            intercept_connection.ignore_code()
    except:
        e = sys.exc_info()[0]
        print("Closing connection: ", e)
        intercept_connection.close()


def command():
    command_connection = pydsfapi.CommandConnection(debug=False)
    command_connection.connect()

    try:
        print(command_connection.get_serialized_machine_model())
    finally:
        command_connection.close()


def subscribe():
    subscribe_connection = pydsfapi.SubscribeConnection(SubscriptionMode.FULL, debug=True)
    subscribe_connection.connect()

    try:
        machine_model = subscribe_connection.get_machine_model()
        print(machine_model.__dict__)
    finally:
        subscribe_connection.close()


async def respond_something(http_endpoint_connection):
    await http_endpoint_connection.read_request()
    await http_endpoint_connection.send_response(200, "so happy you asked for it!")
    http_endpoint_connection.close()


def custom_endpoint():
    cmd_conn = pydsfapi.CommandConnection(debug=True)
    cmd_conn.connect()
    endpoint = None
    try:
        endpoint = cmd_conn.add_http_endpoint(basecommands.HttpEndpointType.GET, "custom", "getIt")
        endpoint.set_endpoint_handler(respond_something)
        time.sleep(1800)
    finally:
        if endpoint is not None:
            endpoint.close()
        cmd_conn.close()


# intercept()
# command()
# subscribe()
custom_endpoint()  # This can only be run on the SBC
