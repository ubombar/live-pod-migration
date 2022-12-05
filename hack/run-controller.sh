#!/usr/bin/bash

# Run this script as ./hack/run-controller.sh from the project root

ROOT_DIR=$( pwd )

## Assuming this is the kubeconfig file of the minikube. Change this if you are using something different.
MINIKUBE_KUBECONFIG="$HOME/.kube/config"

echo "Vendoring..."
go mod vendor

echo "Building..."
go build -o ${ROOT_DIR}/bin/controller ${ROOT_DIR}/cmd/livepodmigration

echo "Running..."
eval ${ROOT_DIR}/bin/controller --kubeconfig $MINIKUBE_KUBECONFIG