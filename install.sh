#!/bin/bash

sudo apt-get update
sudo apt install -y protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

export PATH="$PATH:$(go env GOPATH)/bin
