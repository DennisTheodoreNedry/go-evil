from os import listdir, mkdir
from os.path import isdir, exists, basename
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
            if ".evil" in file:
                to_return.append(file)
    
    return to_return


def compile_file(file_path:str) -> int:
    """
    Description: Compiles the file provided in file_path
    Input:
        - file_path (str) The path to the file you want to compile
    Return: None
    """

    try:
        print(f"[~] Trying to compile {file_path}...", end='')
        subprocess.check_output(["./gevil", "-f", file_path, "-o", f"./output/{basename(file_path).split('.')[0]}"])
        print(" Successful!")
        
    except subprocess.CalledProcessError as err:
        print(f" Failed {err}")
        exit(1)

    

def main() -> None:
    """
    Description: Main method
    Input: None
    Return: None
    """
    if not exists("./gevil"):
        exit("[!] Failed to find the compiler 'gevil' in the root directory of this project")

    if not exists("./output"):
        mkdir("./output")

    files = gather_files("./examples")

    for file in files:
       compile_file(file)

    print(f"\n# All {len(files)} files succeeded to compile!")


if __name__ == "__main__":
    main()