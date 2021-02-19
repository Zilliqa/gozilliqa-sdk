#!/bin/bash
protoc --proto_path=./protobuf --go_out=./protobuf --go_opt=paths=source_relative message.proto --experimental_allow_proto3_optional=true
