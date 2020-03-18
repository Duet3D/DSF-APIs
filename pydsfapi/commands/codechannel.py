"""
codechannel contains an enum with available code channels.

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
from enum import IntEnum


class CodeChannel(IntEnum):
    """Enumeration of every available code channel"""
    HTTP = 0
    Telnet = 1
    File = 2
    USB = 3
    AUX = 4
    Trigger = 5
    CodeQueue = 6
    LCD = 7
    SPI = 8
    Daemon = 9
    AutoPause = 10

    DEFAULT_CHANNEL = SPI
