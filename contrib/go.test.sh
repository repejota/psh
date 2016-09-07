#!/usr/bin/env bash

set -e
echo "" > cover.out

for d in $(go list ./... | grep -v vendor); do
    go test -v -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> cover.out
        rm profile.out
    fi
done
