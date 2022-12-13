#!/usr/bin/bash

CWD=$( pwd )

eval $CWD/hack/update-protobuff.sh
eval $CWD/hack/update-codegen.sh