#!/bin/bash

cd $HOME/go/src/github.com/iafoosball/livematches-service
protoc -I proto/ proto/livematch.proto --go_out=plugins=grpc:proto
