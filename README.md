# lpm-mpod-controller
[Watch the demo](https://youtu.be/7GHJsDL4Bt0)

## Demo in Command Line
Here is the script used in the demo video.

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

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/lpm-mpod-controller:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/lpm-mpod-controller:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

