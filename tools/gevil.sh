#!/bin/sh

if [ ! -d "./builds" ]; then
    mkdir builds
fi

docker run -v $(pwd)/builds:/app/builds --rm goevil/gevil -bd builds "$@"