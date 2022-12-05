#!/usr/bin/bash

# Run this script as ./hack/run-controller.sh from the project root

ROOT_DIR=$( pwd )

## Assuming this is the kubeconfig file of the minikube. Change this if you are using something different.
MINIKUBE_KUBECONFIG="$HOME/.kube/config"

vendoring="false"

if [[ $vendoring == "true" ]]
then 
    echo "Vendoring..."
    # go mod vendor
fi 


echo "Building..."
go build -o ${ROOT_DIR}/bin/controller ${ROOT_DIR}/cmd/controller

echo "Running..."
eval ${ROOT_DIR}/bin/controller --kubeconfig $MINIKUBE_KUBECONFIG