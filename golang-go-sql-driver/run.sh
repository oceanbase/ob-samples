#!/usr/bin/env bash
cd `dirname $0`
CGO_ENABLED=0 GO111MODULE=on go run Test.go