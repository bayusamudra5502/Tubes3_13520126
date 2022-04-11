#!/bin/bash

mkdir -p bin
cd ./src/backend
go build -o ../../bin/server.app
