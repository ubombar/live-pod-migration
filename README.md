# Live Pod Migration

## How to Test?

You can use `./hack/build` script to build both of the `migratord` daemon and `migctl` utility. 

To test the program in minikube, you just need to use `./hack/start-minikube.sh`. This script will setup the minikube with 2 nodes and enable docker experimental mode. Then you need to copy the executables and run them with respected commands.

## What's Next?

This project aims to implement this functionality in kubernetes. We will se what future will bring.

## About the Author
My name is Ufuk Bombar. Feel free to check my github profile [ubombar](https://github.com/ubombar) or contact me regarding this repository at ufukbombar@gmail.com. 

Live Pod Migration repository is for the PROJECT course I am taking in my master's degree in distributed computing at Sorbonne University. My supervisor for this project is [bsenel](https://github.com/bsenel).