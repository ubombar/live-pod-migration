#!/usr/bin/bash

NODES=( "mentor" "mentor-m02" )
PROFILE="mentor"

./hack/gen-code.sh
./hack/build.sh

for node in ${NODES[@]}
do 
    echo "> Installing migratord on $node"

    node_ip=$( minikube ip -p $PROFILE -n $node )
    node_port="4545"


    minikube ssh -p $PROFILE -n $node -- "rm -f /home/docker/migratord"
    # minikube cp -p $PROFILE ./bin/migratord $node:/home/docker/migratord &> /dev/null
    scp -i "$( minikube ssh-key -p $PROFILE -n $node )" ./bin/migratord docker@$node_ip:/home/docker/migratord 

    echo -n "  > Starting migratord on $node..."

    eval minikube ssh -p $PROFILE -n $node -- "/home/docker/migratord --address $node_ip --port $node_port" &

    echo "done!"
done
