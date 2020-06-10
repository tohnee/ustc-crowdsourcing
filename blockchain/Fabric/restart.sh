#!/bin/bash
rm -rf *.key
make clean
make env-up
go build
./compete-service
