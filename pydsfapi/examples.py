#!/usr/bin/env python3

import sys
import time
from pydsfapi import pydsfapi
from pydsfapi.commands import basecommands, code
from pydsfapi.initmessages.clientinitmessages import InterceptionMode, SubscriptionMode

def intercept():
    """Example of intercepting and interacting with codes"""
    intercept_connection = pydsfapi.InterceptConnection(InterceptionMode.PRE, debug=True)
    intercept_connection.connect()

    try:
        while True:

            # Wait for a code to arrive
            cde = intercept_connection.receive_code()

            # Flush the code's channel to be sure we are being in sync with the machine
            success = intercept_connection.flush(cde.channel)

            # Flushing failed so we need to cancel our code
            if not success:
                print('Flush failed')
                intercept_connection.cancel_code()
                continue

            # Check for the type of the code
            if cde.type == code.CodeType.MCode:

                # Do whatever needs to be done if this is the right code
                print(cde, cde.flags)

            # We here ignore it so it will be continued to be processed
            intercept_connection.ignore_code()
    except:
        e = sys.exc_info()[0]
        print("Closing connection: ", e)
        intercept_connection.close()


def command():
    """Example of a command connection to send arbitrary commands to the machine"""
    command_connection = pydsfapi.CommandConnection(debug=False)
    command_connection.connect()

    try:
        # Perform a simple command and wait for its output
        res = command_connection.perform_simple_code('M115')
        print(res)
    finally:
        command_connection.close()


def subscribe():
    """Example of a subscribe connection to get the machine model"""
    subscribe_connection = pydsfapi.SubscribeConnection(SubscriptionMode.PATCH, debug=True)
    subscribe_connection.connect()

    try:
        # Get the complete model once
        machine_model = subscribe_connection.get_machine_model()
        print(machine_model.__dict__)

        # Get 10 updates
        for i in range(0, 10):
            update = subscribe_connection.get_machine_model_patch()
            print(update)
    finally:
        subscribe_connection.close()


async def respond_something(http_endpoint_connection):
    await http_endpoint_connection.read_request()
    await http_endpoint_connection.send_response(200, "so happy you asked for it!")
    http_endpoint_connection.close()


def custom_endpoint():
    """Example to create a custom GET endpoint at http://duet3/machine/custom/getIt"""
    cmd_conn = pydsfapi.CommandConnection(debug=True)
    cmd_conn.connect()
    endpoint = None
    try:
        # Setup the endpoint
        endpoint = cmd_conn.add_http_endpoint(basecommands.HttpEndpointType.GET, "custom", "getIt")
        # Register our handler to reply on requests
        endpoint.set_endpoint_handler(respond_something)
        # This just simulates doing other things as the above runs async
        time.sleep(1800)
    finally:
        if endpoint is not None:
            endpoint.close()
        cmd_conn.close()


intercept()
# command()
# subscribe()
# custom_endpoint()  # This can only be run on the SBC
