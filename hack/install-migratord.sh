#!/usr/bin/bash

NODES=( "mentor" "mentor-m02" )
NODES=( "mentor" )
PROFILE="mentor"

./hack/gen-code.sh
./hack/build.sh

for node in ${NODES[@]}
do 
    echo "> Installing migratord on $node"

    node_ip=$( minikube ip -p $PROFILE -n $node )
    node_port="4545"

    minikube ssh -p $PROFILE -n $node -- "killall migratord" >/dev/null 2>/dev/null

    minikube ssh -p $PROFILE -n $node -- "rm -f /home/docker/migratord" >/dev/null 2>/dev/null

    scp -i "$( minikube ssh-key -p $PROFILE -n $node )" ./bin/migratord docker@$node_ip:/home/docker/migratord >/dev/null

    echo -n "  > Starting migratord on $node..."

    minikube ssh -p $PROFILE -n $node -- "/home/docker/migratord --address $node_ip --port $node_port" &

    echo "done!"
done

sleep 1
echo "> Installation of migratord complete!"