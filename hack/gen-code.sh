#!/usr/bin/bash

echo "Generating protobuffers"

PROTOC="/usr/bin/protoc"
SRC_DIR=$( pwd )

PROTO_FILE="$SRC_DIR/pkg/proto/adapter.proto"
DST_DIR="$SRC_DIR/pkg/generatedpb"

# Create directory
mkdir -p $DST_DIR

$PROTOC -I=$SRC_DIR --go_out=$DST_DIR $PROTO_FILE