#!/usr/bin/env bash

echo "===> Checking code format with gofmt requirements"
gofmt_files=$(find . -name '*.go' | grep -v vendor | xargs gofumpt -l -s)

if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit
fi

exit 0