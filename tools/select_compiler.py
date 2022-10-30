#!/usr/bin/env python3

from re import findall
from os import listdir, environ
from sys import argv
from subprocess import run

TARGET_DIR = environ["GOBIN"]
GRAB_HIGHEST_VER = "gevil_([0-9]\.[0-9])"

def main() -> None:
    """
    Description: Wrapper method which binds the program togheter
    Input: None
    Return: None
    """
    compilers = {}


    # Finding all compilers
    for folder in listdir(TARGET_DIR):
        result = findall(GRAB_HIGHEST_VER, folder)
        if result != None and result != []:
            compilers[result[0]] = f"{TARGET_DIR}/{folder}"
    
    # Identify the newest compiler by checking the number
    newest_compiler = max(compilers.keys())

    # Construct our command string
    command_string = [compilers[newest_compiler]]
    for arg in argv[1:]:
        command_string.append(arg)

    # Runs the compiler
    run(args=command_string)


if __name__ == "__main__":
    main()