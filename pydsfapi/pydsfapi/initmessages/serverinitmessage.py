"""
serverinitmessage holds everything relevant to the first message received from the server

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


class IncompatibleVersionException(Exception):
    """Exception raised when the server and client are incompatible"""


class ServerInitMessage:
    """
    An instance of this class is sent by the server to the client
    in JSON format once a connection has been established.
    """
    @classmethod
    def from_json(cls, data):
        """Deserialize a dictionary coming from JSON into an instance of this class"""
        return cls(**data)

    PROTOCOL_VERSION = 10

    def __init__(self, version: int, id: int):
        self.version = version
        self.id = id

    def is_compatible(self):
        """Check if the message received from the server indicates compatibility with this client"""
        return self.version >= ServerInitMessage.PROTOCOL_VERSION
