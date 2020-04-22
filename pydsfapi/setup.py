#!usr/bin/env python3

from distutils.core import setup

setup(name='pydsfapi',
      version='2.1.0',
      description='Python interface to access DuetSoftwareFramework',
      author='Manuel Coenen',
      author_email='manuel.coenen@gmail.com',
      url='https://github.com/Duet3D/DSF-APIs',
      packages=[
          'pydsfapi',
          'pydsfapi.commands',
          'pydsfapi.initmessages',
          'pydsfapi.model',
          ],
      license='LGPLv3',
      )
