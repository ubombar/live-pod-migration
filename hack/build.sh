#!/usr/bin/bash

TOBUILD=( "migratord" "migctl" )
cwd=$( pwd )

for file in ${TOBUILD[@]}; do 
    echo -n "Building $file..."
    go build -o $cwd/bin/$file $cwd/cmd/$file/main.go
    echo "done!"
done 