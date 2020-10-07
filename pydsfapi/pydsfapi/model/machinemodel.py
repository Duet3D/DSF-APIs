"""
machinemodel contains a generic implementation for the machine model.

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


class MachineModel(dict):
    """
    MachineModel provides generic access to the machine model.
    """
    @classmethod
    def from_json(cls, data):
        """Deserialize an instance of this class from JSON deserialized dictionary"""
        return cls(**data)

    def __getitem__(self, key):
        """Return the item with key key. Returns None if key is not present"""

        # Simple route
        if '.' not in key:
            return dict.__getitem__(self, key)

        parts = key.split('.')
        item = None
        try:
            item = dict.__getitem__(self, parts.pop(0))
        except KeyError:
            return None
        for part in parts:
            try:
                item = item[part]
            except KeyError:
                return None

