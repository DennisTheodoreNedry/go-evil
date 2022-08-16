#!/usr/bin/bash

compiler=~/go/bin/gomarkdoc
folders=(domains utility)

echo "[!] Assuming that the go folder is located in the current users homedirectory"

for folder in "${folders[@]}"
do
    echo "[*] Generating documentation for" $folder

    for subfolder in $folder/*
    do
        mkdir -p docs/$subfolder
        
        for file in $subfolder/*
        do
            $compiler -u -e $file > docs/$file.md
            break # We only want to generate a markdown/package as they will otherwise contain the same info
        done
    done
done


echo "[!] Done, you can find the result in the folder called 'docs'"