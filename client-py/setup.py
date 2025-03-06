from setuptools import setup, find_packages

setup(
    name="godis",
    version="0.1.1",
    author="Javier Avila",
    author_email="njavilas2015@gmail.com",
    description="Cliente de python para Godis",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown",
    url="https://github.com/njavilas2015/godis",
    packages=find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.6",
    install_requires=[],
)
