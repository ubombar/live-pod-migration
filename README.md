# Live Pod Migration

## Demo in Command Line

To create a simple stateful container, the following command can be used. 


```bash
sudo podman run -dt --name looper busybox /bin/sh -c 'i=0; while true; do echo $i; i=$(expr $i + 1); sleep 1; done'
```
<!-- sudo podman logs -l -->

The checkpoint the pod using this command.

<!-- ```bash
sudo podman container checkpoint -l --export=/tmp/checkpoint.tar.gz
``` -->

<!-- scp /tmp/chkpt.tar.gz <destination-host>:/tmp -->

Then restore the command.

<!-- ```bash
sudo podman container restore --import=/tmp/checkpoint.tar.gz
``` -->

## Demo of Migratord and Migctl

Run the migratord as root. Then use migctl to create a migration.

```bash
migctl job --server-address 192.168.122.19--client-address 192.168.122.92 looper
```



<!-- ```bash -->

<!-- # docker run -id --name test centos /bin/sh 'i=0; while true; do echo $i; i=$(expr $i + 1); sleep 1; done' -->


<!-- # docker run --security-opt=seccomp:unconfined --name cr -d ubuntu /bin/sh -c 'i=0; while true; do echo $i; i=$(expr $i + 1); sleep 1; done' -->
<!-- ``` -->

<!-- Normally a checkpoint can be created via the command `docker checkpoint create cr cr-checkpoint`. -->


<!-- 'docker start --checkpoint cr-checkpoint cr' -->

<!-- ## How to Test?

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
``` -->

## What's Next?

This project aims to implement this functionality in kubernetes. We will se what future will bring.

## About the Author
My name is Ufuk Bombar. Feel free to check my github profile [ubombar](https://github.com/ubombar) or contact me regarding this repository at ufukbombar@gmail.com. 

Live Pod Migration repository is for the PROJECT course I am taking in my master's degree in distributed computing at Sorbonne University. My supervisor for this project is [bsenel](https://github.com/bsenel).