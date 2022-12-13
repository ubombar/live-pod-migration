# Testing

## Testing with minikube on Fedora
The minikube start command didn't work in Fedora 37 whuch is the OS I am using to debug and develop this project. The error message was about `kubelet` crashing due to a cgroup driver. There are issues talking about this problem however, they promote workarounds such as downgrading the kubernetes version. In our case this cannot be done since we are utilizing the ContainerCheckpoint feature gate which is introduced in kubernetes 1.25.0.

You can see the error message below:
```
    
```

Solution: I switched to KVM from docker.