#!/usr/bin/env python3

"""
Example of intercepting and interacting with codes

Make sure when running this script to have access to the DSF UNIX socket owned by the dsf user.
"""

import subprocess

import pydsfapi
from pydsfapi.commands.basecommands import MessageType


def start_intercept():
    intercept_connection = pydsfapi.InterceptConnection(pydsfapi.InterceptionMode.PRE)
    intercept_connection.connect()

    try:
        while True:
            # Wait for a code to arrive
            cde = intercept_connection.receive_code()

            # Check for the type of the code
            if cde.type == pydsfapi.CodeType.MCode and cde.majorNumber == 1234:
                # --------------- BEGIN FLUSH ---------------------
                # Flushing is only necessary if the action below needs to be in sync with the machine
                # at this point in the GCode stream. Otherwise it can an should be skipped

                # Flush the code's channel to be sure we are being in sync with the machine
                success = intercept_connection.flush(cde.channel)

                # Flushing failed so we need to cancel our code
                if not success:
                    print("Flush failed")
                    intercept_connection.cancel_code()
                    continue
                # -------------- END FLUSH ------------------------

                # Do whatever needs to be done if this is the right code
                print(cde, cde.flags)

                # Resolve it so that DCS knows we took care of it
                intercept_connection.resolve_code()
            elif cde.type == pydsfapi.CodeType.MCode and cde.majorNumber == 7722:
                # Do whatever needs to be done if this is the right code
                subprocess.run(["sudo", "shutdown", "+1"])
                # Resolve it with a custom response message text
                intercept_connection.resolve_code(
                    MessageType.Warn, "Shutting down SBC in 1min..."
                )
            else:
                # We did not handle it so we ignore it and it will be continued to be processed
                intercept_connection.ignore_code()
    except Exception as e:
        print("Closing connection: ", e)
        intercept_connection.close()


if __name__ == "__main__":
    start_intercept()
