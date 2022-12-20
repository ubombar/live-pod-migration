#!/usr/bin/bash

./hack/gen-code.sh
./hack/build.sh

echo "Installing migctl"

cd ./bin 

sudo install migctl /usr/local/bin

echo "Installed migctl! Use 'source <(migctl completion [shell])' to add command completion."