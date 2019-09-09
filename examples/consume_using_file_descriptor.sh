#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# This is a file descriptor set generated based on:
# https://github.com/protocolbuffers/protobuf/blob/master/examples/addressbook.proto
# To complile the file descriptor set based on the proto definition file:
# protoc --descriptor_set_out addressbook.fds --include_imports addressbook.proto
FILE_DESCRIPTOR_SET_FILE=$DIR/../pkg/reflection/testdata/addressbook.fds
PROTO_NAME=tutorial.Person
BROKER=kafka:9092
TOPIC=person

kafka-protobuf-console consume --brokers $BROKER --descriptor $FILE_DESCRIPTOR_SET_FILE --name $PROTO_NAME -t $TOPIC -f