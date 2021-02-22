#!/bin/bash

#compile app
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o backend

#start app after wait the database
sh -c "/wait && ./backend"