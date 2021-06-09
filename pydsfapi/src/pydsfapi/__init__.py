SOCKET_DIRECTORY = "/run/dsf"
SOCKET_FILE = "dcs.sock"
FULL_SOCKET_PATH = SOCKET_DIRECTORY + "/" + SOCKET_FILE
DEFAULT_BACKLOG = 4

raise DeprecationWarning(
"""This module was deprecated. There will be no more updates or bug fixes.
Please move to the new module dsf-python.
Further instructions can be found at https://github.com/Duet3D/dsf-python""")
