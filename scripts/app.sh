#!/bin/bash

case $1 in
    "serve")
        go run src/main.go serve
        ;;
    "migrate")
        go run src/main.go $@
        ;;
    *)
        echo "Usage: $0 {serve|migrate}"
        exit 1
        ;;
esac