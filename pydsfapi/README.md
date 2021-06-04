# pydsfapi
Python interface to access DuetSoftwareFramework.

## Installation
This package contains a `setup.py` so it can be installed with `python3 setup.py install`.

## Usage
See included `examples/` folder for various use cases.

## License
All files of `pydsfapi` are licensed under LGPLv3.

## Development
For development install globally `pip install sphinx tox` or locally `pip install --user sphinx tox`.

For further information visit:
- [Tox](https://tox.readthedocs.io/en/latest/)
- [Setuptools](https://pypi.org/project/setuptools/)
- [Pytest](https://pytest.org/) install pytest to run test scripts
- [MyPy](http://www.mypy-lang.org/)
- [Black](https://github.com/psf/black)
- [Flake8](http://flake8.pycqa.org)
- [Sphinx](https://docs.readthedocs.io/en/stable/intro/getting-started-with-sphinx.html)
-
### connect to SBC and froward remote socket to local machine

```sh
$ export DSF_SOCK=/tmp/dsf/dsf.sock; mkdir -p $(dirname $DSF_SOCK) && ssh -N -L $DSF_SOCK:/var/run/dsf/dcs.sock pi@duet3.local
```

### run all tests

```sh
$ pytest tests
```

### run single test

```sh
$ pytest tests/test_name.py
```
