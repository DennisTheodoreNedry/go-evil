from os import listdir
from os.path import isdir, exists
import subprocess
from sys import exit

def gather_files(dir:list) -> list:
    """
    Description: Gathers all the test files from the examples directory
    Input:
        - dir (list) The directory to gather files from
    Return: All found files
    """
    to_return = []

    for file in listdir(dir):
        file = f"{dir}/{file}"

        if isdir(file):
            for d_file in gather_files(file):
                to_return.append(d_file)
        else:
            to_return.append(file)
    
    return to_return


def compile_file(file_path:str) -> int:
    """
    Description: Compiles the file provided in file_path
    Input:
        - file_path (str) The path to the file you want to compile
    Return: -1 if it failed to compile or 1 if it succeeded to compile
    """
    toReturn = 1

    try:
        print(f"[!] Trying to compile {file_path}...", end='')
        subprocess.check_output(["./gevil", "-f", file_path])
        print(" Successful!")
    except subprocess.CalledProcessError as err:
        print(f" Failed {err}")
        toReturn = -1

    return toReturn
    

def main() -> None:
    """
    Description: Main method
    Input: None
    Return: None
    """
    if not exists("./gevil"):
        exit("Failed to find the compiler 'gevil' in the root directory of this project")

    files = gather_files("./examples")
    status = 0

    for file in files:
        status += compile_file(file)

    print(f"\n# Result: {status}/{len(files)} succeeded to compile!")


if __name__ == "__main__":
    main()