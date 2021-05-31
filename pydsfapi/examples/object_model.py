#!/usr/bin/env python3

from pydsfapi.connections import CommandConnection

class Mcode:
    def __init__(self, code, param=False, timeout=4, desc="missing"):
        self.code=code
        self.param = param
        self.desc=desc

    def print(self):
        print("M{code} '{desc}' ".format(code=self.code, desc=self.desc))


    def send(self, conn):
        cmd = "M{} {}".format(self.code, self.param)
        print("cmd {}".format(cmd))
        return conn.perform_simple_code(cmd)

mcodes = [
    Mcode(code=115, desc="Print firmware version or set hardware type"),
    Mcode(code=409, param="F\"v\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"boards\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"directories\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"fans\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"global\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"heat\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"inputs\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"job\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"limits\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"move\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"network\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"scanner\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"sensors\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"seqs\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"spindles\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"state\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"tools\"", desc="Get object model values in JSON format"),
    Mcode(code=409, param="K\"volumes\"", desc="Get object model values in JSON format"),
]


def send_all():
    conn = CommandConnection(debug=False)
    conn.connect()
    print("opened connection")

    for entry in mcodes:
        try:
            entry.send(conn)
        except TimeoutError as error:
            print("Warning: {}".format(error))
            break
        except Exception as error:
            print("Error: {}".format(error))
            break

    conn.close()
    print("closed connection")


if __name__ == "__main__":
    send_all()
