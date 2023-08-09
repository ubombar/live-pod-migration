#!/usr/bin/bash

# Run this script as ./hack/run-migrator.sh from the project root

ROOT_DIR=$( pwd )

## Assuming this is the kubeconfig file of the minikube. Change this if you are using something different.
MINIKUBE_KUBECONFIG="$HOME/.kube/config"

vendoring="false"

if [[ $vendoring == "true" ]]
then 
    echo "Vendoring..."
    go mod vendor
fi 


echo "Building..."
go build -o ${ROOT_DIR}/bin/migrator ${ROOT_DIR}/cmd/migrator

echo "Running..."
eval ${ROOT_DIR}/bin/migrator --kubeconfig $MINIKUBE_KUBECONFIG