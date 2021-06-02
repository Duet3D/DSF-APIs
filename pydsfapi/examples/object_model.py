#!/usr/bin/env python3

from pydsfapi.connections import CommandConnection

import re
import logging


class TestMcode:
    def __init__(self, code, param=False, result=None, timeout=4, desc="missing"):
        self.code = code
        self.param = param
        self.result = result
        self.desc = desc

    def print(self):
        logging.info("M{code} '{desc}' ".format(code=self.code, desc=self.desc))

    def getCommand(self):
        return "M{} {}".format(self.code, self.param)

    def test(self, result):
        if self.result is None:
            logging.info("M{} {} result: {}".format(self.code, self.param, result))
            return False

        logging.debug("testing\n  {} {}\n  {} {}\n".format(
            type(result), result, type(self.result), self.result
        ))
        res = re.match(self.result, result)
        return res is not None


    def send(self, conn):
        return conn.perform_simple_code(self.getCommand())

    def other(self, conn):
        pass


mcodes = [
    TestMcode(code=115, result='FIRMWARE_NAME: RepRapFirmware for Duet 3 Mini 5\+ FIRMWARE_VERSION: .+ ELECTRONICS: .+ FIRMWARE_DATE: .+', desc="Print firmware version or set hardware type"),
    TestMcode(code=409, param="F\"v\"", desc="Get object model values in JSON format", result='{"key":"","flags":"v","result":{"boards":\[{}\],"directories":{},"fans":\[{},{},{}\],"global":{},"heat":{},"inputs":\[{},{},{},{},{},{},{},{},{},{},{},{}\],"job":{},"limits":{},"move":{},"network":{},"scanner":{},"sensors":{},"seqs":{},"spindles":\[{},{}\],"state":{},"tools":\[{}\],"volumes":\[{},{}\]}}'),
    TestMcode(code=409, param="K\"boards\" F\"d99vn\"", desc="Get object model values in JSON format", result='{"key":"boards","flags":"d99vn","result":\[{"canAddress":\d+,"directDisplay":\w+,"firmwareDate":"[\d-]+","firmwareFileName":"[\w_\.]+".*,"firmwareName":"[\w_ \+\.]+","firmwareVersion":"[\w\-_\.]+","iapFileNameSBC":"[\w\-_\.]+","iapFileNameSD":"[\w\-_\.]+","maxHeaters":\d+,"maxMotors":\d+,"mcuTemp":{"current":\d+.\d+,"max":\d+.\d+,"min":\d+.\d+},"name":"[\w \-_\+\.]+","shortName":"[\w \-_\+\.]+","supportsDirectDisplay":(true|false),"uniqueId":"[\w-]+","vIn":{"current":\d+.\d+,"max":\d+.\d+,"min":\d+.\d+}}\],"next":\d+}'),
    # TestMcode(code=409, param="K\"boards\" F\"d2vn\"", desc="Get object model values in JSON format"),
    TestMcode(code=409, param="K\"directories\" F\"d99vn\"", desc="Get object model values in JSON format", result='{"key":"directories","flags":"d99vn","result":{"filaments":"\d:/filaments/","firmware":"\d:/firmware/","gCodes":"\d:/gcodes/","macros":"\d:/macros/","menu":"\d:/menu/","scans":"\d:/scans/","system":"\d:/sys/","web":"\d:/www/"}}'),
    # TestMcode(code=409, param="K\"fans\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"global\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"heat\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"inputs\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"job\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"limits\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"move\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"network\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"scanner\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"sensors\" F\"d99vn\"", desc="Get object model values in JSON format"),
    TestMcode(code=409, param="K\"seqs\" F\"d99vn\"", result='{"key":"seqs","flags":"d99vn","result":{"boards":0,"directories":0,"fans":\d+,"global":0,"heat":\d+,"inputs":\d+,"job":\d+,"move":\d+,"network":\d,"reply":\d+,"scanner":\d+,"sensors":\d+,"spindles":\d+,"state":\d+,"tools":\d+,"volChanges":\[\d+,\d+\],"volumes":\d+}}', desc="Get object model values in JSON format"),
    TestMcode(code=409, param="K\"spindles\" F\"d99vn\"", result='{"key":"spindles","flags":"d99vn","result":\[{"active":\d+,"current":\d+,"frequency":\d+,"max":\d+,"min":\d+,"state":"unconfigured"},{"active":\d+,"current":\d+,"frequency":\d+,"max":\d+,"min":\d+,"state":"unconfigured"}\],"next":\d+}', desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"state\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"tools\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"volumes\" F\"d99vn\"", desc="Get object model values in JSON format"),
    # TestMcode(code=409, param="K\"invalid\" F\"d99vn\"", result='{"key":"invalid","flags":"d99vn","result":null}', desc="Get object model values in JSON format"),
]


def send_all():
    conn = CommandConnection(debug=False)
    conn.connect()
    print("opened connection")

    result = None

    for entry in mcodes:
        try:
            result = entry.send(conn)
        except TimeoutError as error:
            logging.warning("Warning: {}".format(error))
            break
        except Exception as error:
            logging.error("Error: {}".format(error))
            break

        if (entry.test(result)):
            logging.info("{}: OK".format(entry.getCommand()))
        else:
            logging.error("{}: FAIL".format(entry.getCommand()))

    conn.close()
    print("closed connection")


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.DEBUG,
        format='*** %(asctime)s - %(message)s',
        filename=None
    )

    logging.info("starting tests")
    send_all()
