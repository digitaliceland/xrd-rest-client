#!/bin/bash

echo compling for Windows: release/windows/rest-client.exe
mkdir -p release/windows
GOOS=windows GOARCH=amd64 go build -o release/windows/rest-client.exe main.go

echo compling for MacOS: release/macos/rest-client
mkdir -p release/macos
GOOS=darwin GOARCH=amd64 go build -o release/macos/rest-client main.go

echo compling for Linux: release/linux/rest-client
mkdir -p release/linux
GOOS=linux GOARCH=amd64 go build -o release/linux/rest-client main.go

