# Possible problems and responses

### How to access kubelet storage?
You can access the host node's storage by mounting node's path to pod.

### Where the kubernetes checkpoints are stored?
They are in `/var/lib/kubelet/checkpoints/*.zip`.

### How to access node's container runtime?
For now I will try to access it by `/run/containerd/containerd.sock`. This bypasses the *kubelet* which can cause many problems.

### How to create a development environment
Testing the code in a cluster is actually time consuming and tricky. After some reseach I decided on the following idea. Please note that I am not an expert on container and cluster technologies and this is a proof of concept.
* First, I will compile the go code in my local machine, then I will have the migrator and controller executables. Here I can use the controller executable as it is with giving the kubeconfig file path. However, it is a bit tricky to setup the migrator since it needs to communicate with the kubelet and needs to be on the actual node.
* Thus, for the migrator, I will create a docker image for kubernetes to use. However, it turns out the way I think would work. My idea was to point the docker image registry to my local docker registry (on my local machine instead of KVM nodes). So, I need to create a private local registry in my local machine.
* After I do this I can simply push the migrator image I created to the registry. And let kubernetes use that to pull the images. However, it takes quite a long time to setup all this and to make a small change I need to rebuild the image push it and pull it on the kubernetes then redeploy the deamonsets. There might be a better method but for now I will use this.
> This might be different for everyone but my docker registry is on `192.168.39.1` and the network of `host`, `mentor` and `mentor-m02` is `192.168.39.1/24`.