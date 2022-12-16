#!/usr/bin/bash

PROFILE="mentor"

echo "> Checking if the cluster is up..."

running=$( minikube status -p ${PROFILE} | grep Running | wc -l )

expmode='{  \"exec-opts\": [    \"native.cgroupdriver=systemd\"  ],  \"log-driver\": \"json-file\",  \"log-opts\": {    \"max-size\": \"100m\"  },  \"experimental\":true,  \"storage-driver\": \"overlay2\"}'

if [[ $running -lt 1 ]]
then 
    minikube start -p ${PROFILE} --feature-gates=ContainerCheckpoint=true,LocalStorageCapacityIsolation=true
    echo "  > Cluster is now running."
else 
    echo "  > Cluster is already up and running."
fi 

echo "> Checking number of nodes in the cluster..."

# Check if there are at least two nodes
nodecount=$( minikube node list -p ${PROFILE} | wc -l )

if [[ $nodecount -lt 2 ]]
then 
    echo =n "  > There are not enough nodes, adding 1 node to the minikube cluster..."
    minikube node add -p ${PROFILE}
    echo "done!"
else 
    echo "  > There are $nodecount nodes."
fi

NODES=( $( minikube node list -p $PROFILE | grep -o '^\S*' | grep "" ))
# echo "!> You need to enable docker experimental mode for each node to checkpoint and restore to work!"

echo "> Checking if checkpoint restore enabled in nodes..."
intervention=0

for node in ${NODES[@]}
do 
    enabled=$( minikube ssh -p $PROFILE -n $node "docker version -f '{{.Server.Experimental}}'" | grep "true" )

    if ! [ $enabled ]; then 
        intervention=1
        echo -n "  > Experimental mode is disabled on node '$node'"

        minikube ssh -p $PROFILE -n $node "su - root -c \"echo '$expmode' > /etc/docker/daemon.json\""

        minikube ssh -p $PROFILE -n $node "su - root -c \"systemctl restart docker\""

        echo "done!"
    fi
done

if [ $intervention -eq 0 ]; then 
    echo "  > All nodes has experimental mode enabled."
fi

echo "> All clear!"