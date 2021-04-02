#!/usr/bin/env python3

"""
Example to create a custom GET endpoint at http://duet3/machine/custom/getIt

Make sure when running this script to have access to the DSF UNIX socket owned by the dsf user.
"""

import time

import pydsfapi
from pydsfapi.commands.basecommands import HttpEndpointType
from pydsfapi.http_endpoint import HttpEndpointConnection


async def respond_something(http_endpoint_connection: HttpEndpointConnection):
    await http_endpoint_connection.read_request()
    await http_endpoint_connection.send_response(200, "so happy you asked for it!")
    http_endpoint_connection.close()


def custom_http_endpoint():
    cmd_conn = pydsfapi.CommandConnection()
    cmd_conn.connect()
    endpoint = None
    try:
        # Setup the endpoint
        endpoint = cmd_conn.add_http_endpoint(HttpEndpointType.GET, "custom", "getIt")
        # Register our handler to reply on requests
        endpoint.set_endpoint_handler(respond_something)
        # This just simulates doing other things as the above runs async
        time.sleep(1800)
    finally:
        if endpoint is not None:
            endpoint.close()
        cmd_conn.close()


if __name__ == "__main__":
    custom_http_endpoint()
