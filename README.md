# Live Pod Migration

Live Pod Migration aims to migrate a pod running in a Kubernetes cluster from one node to another one while preserving the state of the pod's containers. 

In this project, I am using Kubernetes' implementation of checkpointing API. However, we don't have the restore functionality yet. This is why I am bypassing the kubelet to access the container runtime directly. But don't forget this is an MVP!

## Design
This is an actively changing project. I am also preparing a short documentation that can be accessible in the docs folder. To see how the LPM system is designed click [here](docs/steps-of-the-migration.md).

## Problems
There are many problems with migrating processes if not pods. [Here](docs/problems.md) I am taking notes about possible problems and their solutions. Note that I am not a Kubernetes expert.

## Setting up Cluster
Use the following command to setup a node in the cluster.
```source <(curl -s https://gist.githubusercontent.com/ubombar/0a64ff40a15bcbc5988e23dd28a9ecca/raw/44f2f168452dc44a3f6b35aeaf6b5df3960491ca/kubernetes-installer.sh)```

## About the Author
My name is Ufuk Bombar. Feel free to check my github profile [ubombar](https://github.com/ubombar) or contact me regarding this repository at ufukbombar@gmail.com. 

Live Pod Migration repository is for the PROJECT course I am taking in my master's degree in distributed computing at Sorbonne University. My supervisor for this project is [bsenel](https://github.com/bsenel).
