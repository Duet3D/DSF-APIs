import threading
import os
import pathlib
import socket
import time
import importlib.util

here = pathlib.Path(__file__).parent.parent.resolve()
example_path = here / "examples/custom_m_codes.py"
spec = importlib.util.spec_from_file_location("custom_m_codes", example_path)
custom_m_codes = importlib.util.module_from_spec(spec)
spec.loader.exec_module(custom_m_codes)


def test_custom_m_codes(monkeypatch, tmp_path):
    mock_dcs_socket_path = os.path.join(tmp_path, "dsf.socket")
    monkeypatch.setattr(
        "pydsfapi.connections.InterceptConnection.connect.__defaults__",
        (mock_dcs_socket_path,),
    )

    dcs_passed = threading.Event()

    def mock_dcs():
        server = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        server.bind(mock_dcs_socket_path)
        server.listen(1)
        conn, _ = server.accept()
        conn.sendall(b'{"version": 11, "id": "foobar"}')
        assert (
            conn.recv(1024) == b"{"
            b'"mode":"Intercept","version": 11,"InterceptionMode":"Pre","Channels":["HTTP","Telnet",'
            b'"File","USB","Aux","Trigger","Queue","LCD","SBC","Daemon","Aux2","AutoPause","Unknown"],'
            b'"Filters":null,"PriorityCodes":false'
            b"}"
        )
        conn.sendall(b'{"success":true}')
        conn.sendall(
            b"{"
            b'"connection":{"id":12,"apiVersion":10,"isConnected":true},"sourceConnection":12,'
            b'"result":null,"type":"M","channel":"HTTP","lineNumber":null,"indent":0,"keyword":0,'
            b'"keywordArgument":null,"majorNumber":1234,"minorNumber":null,"flags":2048,"comment":null,'
            b'"filePosition":null,"length":6,"parameters":[],"command":"Code"'
            b"}"
        )
        assert conn.recv(1024) == b'{"command":"Flush","Channel":"HTTP"}'
        conn.sendall(b'{"result":true,"success":true}')
        assert conn.recv(1024) == b'{"command":"Resolve","Type":0,"Content":null}'
        conn.sendall(
            b"{"
            b'"connection":{"id":12,"apiVersion":10,"isConnected":true},"sourceConnection":12,'
            b'"result":null,"type":"M","channel":"HTTP","lineNumber":null,"indent":0,"keyword":0,'
            b'"keywordArgument":null,"majorNumber":5678,"minorNumber":null,"flags":2048,"comment":null,'
            b'"filePosition":null,"length":6,"parameters":[],"command":"Code"'
            b"}"
        )
        assert conn.recv(1024) == b'{"command":"Resolve","Type":0,"Content":null}'
        conn.close()
        dcs_passed.set()  # indicate that all asserts passed and the mock_dcs is shutting down

    thread = threading.Thread(target=mock_dcs, daemon=True)
    thread.start()
    time.sleep(1)

    custom_m_codes.start_intercept()

    thread.join()

    assert dcs_passed.is_set()
