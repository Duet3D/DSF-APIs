#!usr/bin/env python3

from setuptools import setup, find_packages
import pathlib

here = pathlib.Path(__file__).parent.resolve()
long_description = (here / "README.md").read_text(encoding="utf-8")

setup(
    name="pydsfapi",
    version="3.2.0",
    description="Python interface to access DuetSoftwareFramework",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/Duet3D/DSF-APIs",
    author="Manuel Coenen",
    author_email="manuel.coenen@gmail.com",
    classifiers=[
        "Development Status :: 3 - Production/Stable",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Libraries",
        "License :: OSI Approved :: GNU Lesser General Public License v3 (LGPLv3)",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.5",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3 :: Only",
    ],
    keywords="Duet3D, DuetSoftwareFramework, DSF",
    package_dir={"": "src"},
    packages=find_packages(where="src"),
    python_requires=">=3.5, <4",
    extras_require={
        "dev": [
            "sphinx",
        ],
    },
    project_urls={
        "Duet3D Support": "https://forum.duet3d.com/",
        "Bug Reports": "https://github.com/Duet3D/DSF-APIs/issues",
        "Source": "https://github.com/Duet3D/DSF-APIs/",
    },
)
