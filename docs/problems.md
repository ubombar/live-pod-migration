# Possible problems and responses

### How to access kubelet storage?
You can access the host node's storage by mounting node's path to pod.

### Where the kubernetes checkpoints are stored?
They are in `/var/lib/kubelet/checkpoints/*.zip`.

### How to access node's container runtime?
For now I will try to access it by `/run/containerd/containerd.sock`. This bypasses the *kubelet* which can cause many problems.