#!/usr/bin/bash

echo "Generating protobuffers"

PROTOC="/usr/bin/protoc"
SRC_DIR=$( pwd )
PROTO_FILE="$SRC_DIR/pkg/migrator/v1alpha1/migrator_interface.proto"
DST_DIR="$SRC_DIR/pkg/generated/migrator/v1alpha1"

# Create directory
mkdir -p $DST_DIR

$PROTOC -I=$SRC_DIR --go_out=$DST_DIR $PROTO_FILE