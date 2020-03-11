"""
codeparameter contains all classes and methods dealing with deserialized code parameters.

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
import json


class CodeParserException(Exception):
    """Raised if codes could not be parsed properly"""


class CodeParameter(json.JSONEncoder):
    """Represents a parsed parameter of a G/M/T-code"""
    def default(self, o):
        return {'letter': o.letter, 'value': o.value, 'isString': isinstance(o.value, str)}

    @classmethod
    def from_json(cls, data):
        """Instantiate a new instance of this class from JSON deserialized dictionary"""
        return cls(**data)

    @classmethod
    def simple_param(cls, letter: str, value):
        """Create a new simple parameter without parsing the value"""
        return cls(letter, value)

    def __init__(self, letter: chr, value: str, isString: bool = None):
        """
        Creates a new CodeParameter instance and parses value to a native data type
        if applicable
        """

        if isString is None:
            self.letter = letter
            self.string_value = str(value)
            self.__parsed_value = value
            return

        self.letter = letter
        self.string_value = value
        self.is_string = isString
        self.is_expression = False
        self.is_driver_id = False
        if self.is_string:
            self.__parsed_value = value
            return

        value = value.strip()
        # Empty parameters are repesented as integers with the value 0 (e.g. G92 XY => G92 X0 Y0)
        if not value:
            self.__parsed_value = 0
        elif value.startswith('{}') and value.endswith('}'):  # It is an expression
            self.is_expression = True
        elif ':' in value:  # It is an array (or a string)
            split = value.split(':')
            try:
                if '.' in value:  # If there is a dot anywhere, attempt to parse it as a float array
                    self.__parsed_value = list(map(float, split))
                else:  # If there is no dot, it could be an integer array
                    self.__parsed_value = list(map(int, split))
            except:
                self.__parsed_value = value
        else:
            try:
                self.__parsed_value = int(value)
            except:
                try:
                    self.__parsed_value = float(value)
                except:
                    self.__parsed_value = value

    def convert_driver_ids(self):
        """Convert this parameter to driver id(s)"""
        if self.is_expression:
            return

        drivers = []
        parameters = self.string_value.split(':')
        for value in parameters:
            segments = value.split('.')
            segment_count = len(segments)
            if segment_count == 1:
                drivers.append(int(segments[0]))
            elif segment_count == 2:
                driver = (int(segments[0]) << 16) & 0xFFFF
                driver |= (int(segments[1] & 0xFFFF))
            else:
                raise CodeParserException('Driver value from {0} parameter is invalid'.format(
                    self.letter))

        if len(drivers) == 1:
            self.__parsed_value = drivers[0]
        else:
            self.__parsed_value = drivers
        self.is_driver_id = True

    def as_float(self):
        """Conversion to float"""
        if isinstance(self.__parsed_value, float):
            return self.__parsed_value
        if isinstance(self.__parsed_value, int):
            return float(self.__parsed_value)

        raise Exception('Cannot convert {0} parameter to float (value {1})'.format(
            self.letter, self.string_value))

    def as_int(self):
        """Conversion to int"""
        if isinstance(self.__parsed_value, int):
            return self.__parsed_value

        raise Exception('Cannot convert {0} parameter to int (value {1})'.format(
            self.letter, self.string_value))

    def as_float_array(self):
        """Conversion to float array"""
        try:
            if isinstance(self.__parsed_value, list):
                return list(map(float, self.__parsed_value))
            if isinstance(self.__parsed_value, float):
                return [self.__parsed_value]
            if isinstance(self.__parsed_value, int):
                return [float(self.__parsed_value)]
        except:
            pass
        raise Exception('Cannot convert {0} parameter to float array (value {1})'.format(
            self.letter, self.string_value))

    def as_int_array(self):
        """Conversion to int array"""
        try:
            if isinstance(self.__parsed_value, list):
                return list(map(int, self.__parsed_value))
            if isinstance(self.__parsed_value, int):
                return [float(self.__parsed_value)]
        except:
            pass
        raise Exception('Cannot convert {0} parameter to float array (value {1})'.format(
            self.letter, self.string_value))

    def as_bool(self):
        """Conversion to bool"""
        try:
            return float(self.string_value) > 0
        except:
            return False

    def __eq__(self, other):
        if self is None:
            return other is None
        if isinstance(other, CodeParameter):
            return self.letter == other.letter and self.__parsed_value == other.__parsed_value
        return self.__parsed_value == other

    def __ne__(self, other):
        return not self == other

    def __str__(self):
        if self.is_string:
            double_quoted = self.string_value.replace('"', '""')
            return '{0}"{1}"'.format(self.letter, double_quoted)

        return '{0}{1}'.format(self.letter, self.string_value)
