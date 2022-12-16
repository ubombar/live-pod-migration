#!/usr/bin/bash

echo -n "Generating gRPC stub..."

PROTOC="/usr/bin/protoc"
SRC_DIR=$( pwd )

PROTO_PATH="./pkg"
PROTO_FILE="$PROTO_PATH/migrator.proto"
DST_DIR="$SRC_DIR/pkg/migrator"

# Create directory
mkdir -p $DST_DIR

$PROTOC --proto_path=$PROTO_PATH --go_out=$DST_DIR --go-grpc_out=$DST_DIR $PROTO_FILE

echo "done!"
