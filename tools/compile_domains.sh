#!/bin/sh

for i in domains/* ; do
  if [ -d "$i" ]; then
    cd $i && go build -buildmode=plugin && cd ../..
  fi
done