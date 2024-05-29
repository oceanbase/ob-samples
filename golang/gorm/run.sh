#!/usr/bin/env bash
cd `dirname $0`
CGO_ENABLED=0 GO111MODULE=on go run example.go -host 127.0.0.1 -port 3306 -username user@tenant#cluster -password pwd -database dbname
