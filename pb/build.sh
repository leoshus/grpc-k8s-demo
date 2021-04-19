#!/bin/sh

protoc --go_out=plugins=grpc:. ./hello.proto