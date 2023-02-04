#!/usr/bin/bash

PROFILE="mentor"
NODE="mentor"
CONTAINER_NAME="stateful-container"
CONTAINER_PORT="54321"

echo -n "> Removing old container..."

# First stop and remove the old counter
minikube ssh -p $PROFILE -n $NODE -- "docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME" &> /dev/null

echo " done!"

CONTAINER_HASH=$( minikube ssh -p $PROFILE -n $NODE -- "docker run -itd --name $CONTAINER_NAME -ePORT=54321 -p $CONTAINER_PORT:54321 drosenbauer/docker-counter:latest" )

echo "> Container created: $CONTAINER_HASH"

echo $( minikube ssh -p $PROFILE -n $NODE -- curl )

# docker container inspect 