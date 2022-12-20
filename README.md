# Live Pod Migration

## How to Test?

You can use `./hack/build` script to build both of the `migratord` daemon and `migctl` utility. 

To test the program in minikube, you just need to use `./hack/start-minikube.sh`. This script will setup the minikube with 2 nodes and enable docker experimental mode. Then you need to copy the executables and run them with respected commands.

## Demo!
After you create the minikube cluster using the script you can use install scripts to install the `migctl` and `migratord` programs.

```bash
    # Use this script to build and install migctl
    ./hack/install-migctl.sh

    # Use this script to build and install migratword
    ./hack/install-migratord.sh
```

Then see the containers and images

```bash
    # See the docker containers
    minikube ssh -p mentor -n mentor docker ps

    # See the socker images
    minikube ssh -p mentor -n mentor docker images
```

Now create a container.

```bash
    minikube ssh -p mentor -n mentor -- "docker run -itd -ePORT=54321 -p 54321:54321 drosenbauer/docker-counter:latest"
```

Now, we create the migration.

```bash
    migctl job --address-client $( minikube ip -p mentor -n mentor ) --address-server $( minikube ip -p mentor -n mentor-m02 ) --port-client 4545 --port-server 4545 --key $( minikube ssh-key -p mentor -n mentor-m02 ) container_id
```

This command will create a migration job and invoke the migratord. You can watch the migration with the following command.

```bash
    watch -n 1 migctl get --address-client $( minikube ip -p mentor -n mentor ) --address-server $( minikube ip -p mentor -n mentor-m02 ) --port-client 4545 --port-server 4545 migration_id
```

## What's Next?

This project aims to implement this functionality in kubernetes. We will se what future will bring.

## About the Author
My name is Ufuk Bombar. Feel free to check my github profile [ubombar](https://github.com/ubombar) or contact me regarding this repository at ufukbombar@gmail.com. 

Live Pod Migration repository is for the PROJECT course I am taking in my master's degree in distributed computing at Sorbonne University. My supervisor for this project is [bsenel](https://github.com/bsenel).