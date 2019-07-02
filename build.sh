#!/bin/sh
GOOS=linux GOARCH=arm64 go build -ldflags '-s -w' -o ippush
upx ippush;