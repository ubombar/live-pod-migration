#!/usr/bin/bash

NODES=( "mentor" "mentor-m02" )
PROFILE="mentor"



# $( minikube ssh-key -p $PROFILE -n $NODES[0] )

# minikube scp -p mentor 

# minikube cp minikube-m01:a.txt minikube-m02:/home/docker/b.txt