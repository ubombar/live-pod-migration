#!/usr/bin/bash

PROFILE="mentor"

echo "Checking if the cluster is up..."

running=$( minikube status -p mentor | grep Running | wc -l )

if [[ $running -lt 1 ]]
then 
    minikube start -p ${PROFILE} --feature-gates=ContainerCheckpoint=true,LocalStorageCapacityIsolation=true
    echo "> Cluster is now running."
else 
    echo "> Cluster is already up and running."
fi 

echo "Checking number of nodes in the cluster..."

# Check if there are at least two nodes
nodecount=$( minikube node list -p ${PROFILE} | wc -l )

if [[ $nodecount -lt 2 ]]
then 
    echo "> There are not enough nodes, adding 1 node to the minikube cluster"
    minikube node add -p ${PROFILE}
    echo "> Done!"
else 
    echo "> There are $nodecount nodes. You are ready to hack!"
fi
