#!/usr/bin/bash

NODES=( "mentor" "mentor-m02" )
PROFILE="mentor"

./hack/gen-code.sh
./hack/build.sh

for node in ${NODES[@]}
do 
    echo -n "> Installing migratord on $node..."

    node_ip=$( minikube ip -p $PROFILE -n $node )
    node_port="4545"

    minikube ssh -p $PROFILE -n $node -- "killall migratord" >/dev/null 2>/dev/null

    minikube ssh -p $PROFILE -n $node -- "rm -f /home/docker/migratord" >/dev/null 2>/dev/null

    scp -i "$( minikube ssh-key -p $PROFILE -n $node )" ./bin/migratord docker@$node_ip:/home/docker/migratord >/dev/null

    minikube ssh -p $PROFILE -n $node -- "/home/docker/migratord --address $node_ip --port $node_port 1>logs.txt 2>logs.txt" >/dev/null 2>/dev/null &

    echo "done!"
done

echo "> Installation of migratord complete!"