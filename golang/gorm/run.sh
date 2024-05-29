#!/usr/bin/env bash
cd `dirname $0`
CGO_ENABLED=0 GO111MODULE=on go run example.go -host 127.0.0.1 -port 2881 -username root@test -database test
