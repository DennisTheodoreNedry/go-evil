# Creates a list with all current examples

import os
from itertools import chain

def list_files(dir, parent):
    toReturn = []
    for child in dir:
        child = parent+child
        if os.path.isdir(child):
            output = list_files(os.listdir(child), child+"/")
            if len(output) > 0:
                toReturn.append(output)
        else:
            if child.endswith(".ge"): # The only thing we care about
                toReturn.append(child)
    return toReturn

def write_to_file(all_examples):
    with open("examples/examples.txt", "w") as out:
        for line in all_examples:
            out.write(line+"\n")

def algorithm(input_list):
    while True:
        found_list = False
        for i,x in enumerate(input_list):
            if type(x) == list:
                temp = [y for y in x] # Grab every value (this might be a string or a list)
                input_list.pop(i)
                for y in temp:
                    input_list.append(y)
                found_list = True

        if not found_list:
            break
    return input_list


def main():
    high_lvl_dir = os.listdir("examples/")
    high_lvl_dir.remove("generate_list_of_examples.py")
    all_examples_found = list_files(high_lvl_dir, "examples/")
    final = algorithm(all_examples_found)
    write_to_file(final)

if __name__ == "__main__":
    if os.path.exists("examples/examples.txt"): # We need to start clean
        os.remove("examples/examples.txt")
    main()