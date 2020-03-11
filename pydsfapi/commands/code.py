"""
Module code contains relevant classes and enums being received
as well as sent back to the server.

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
from enum import Enum, IntEnum
from .basecommands import BaseCommand
from .codeparameter import CodeParameter
from .result import Message


class CodeType(str, Enum):
    """Type of a generic G/M/T-code. If none is applicable, it is treated as a comment"""
    Comment = 'C'
    GCode = 'G'
    MCode = 'M'
    TCode = 'T'


class Keyword(IntEnum):
    """Types of conditional G-code"""
    KeywordNone = 0
    If = 1
    ElseIf = 2
    Else = 3
    While = 4
    Break = 5
    Return = 6
    Abort = 7
    Var = 8
    Set = 9
    Echo = 10


class CodeFlags(IntEnum):
    """Code bits to classify G/M/T-codes"""
    CodeFlagsNone = 0
    Asynchronous = 1
    IsPreProcessed = 2
    IsPostProcessed = 4
    IsFromMacro = 8
    IsNestedMacro = 16
    IsFromConfig = 32
    IsFromConfigOverride = 64
    EnforceAbsolutePosition = 128
    IsPrioritized = 256
    Unbuffered = 512


class Code(BaseCommand):
    """A parsed representation of a generic G/M/T-code"""
    @classmethod
    def from_json(cls, data):
        """Deserialize an instance of this class from JSON deserialized dictionary"""
        data['result'] = [] if data['result'] is None else list(
            map(Message.from_json, data['result']))
        data['parameters'] = list(map(CodeParameter.from_json, data['parameters']))
        return cls(**data)

    def parameter(self, letter: str, default=None):
        """Retrieve the parameter whose letter equals c or generate a default parameter"""
        letter = letter.upper()
        param = [param for param in self.parameters if param.letter.upper() == letter]
        if len(param) > 0:
            return param[0]
        if default is not None:
            return CodeParameter.simple_param(letter, default)
        return None

    def get_unprecedented_string(self, quote: bool = False):
        """
        Reconstruct an unprecedented string from the parameter list or
        retrieve the parameter which does not have a letter assigned.
        """
        str_list = []
        for param in self.parameters:
            if quote and param.is_string:
                str_list.append(f'{param.letter}"{param.string_value}"')
            else:
                str_list.append(f'{param.letter}{param.string_value}')
        return ' '.join(str_list)

    def __str__(self):
        """Convert the parsed code back to a text-based G/M/T-code"""
        if self.type == CodeType.Comment:
            return f';{self.comment}'

        str_list = []
        str_list.append(self.short_str())

        for param in self.parameters:
            str_list.append(f' {param}')

        if self.comment:
            str_list.append(f' ;{self.comment}')

        if len(self.result) > 0:
            str_list.append(f' => {self.result}')

        return ''.join(str_list)

    def short_str(self):
        """Convert only the command portion to a text-based G/M/T-code (e.g. G28)"""
        if self.type == CodeType.Comment:
            return '(comment)'

        if self.majorNumber is not None:
            if self.minorNumber is not None:
                return f'{self.type}{self.majorNumber}.{self.minorNumber}'

            return f'{self.type}{self.majorNumber}'

        return f'{self.type}'
