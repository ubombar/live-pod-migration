# Steps of the Migration
This project contains two binary files that will be deployed to the cluster. The first one is the `custom controller` and the second one is the `migrator`. The CRD for managing this operation is named `LivePodMigration`. Here is the inside of the CRD:

```yaml
spec:
    # Defines which namespace the pod is
    podNamespace: string
    # Defines the name of the pod
    podName: string
    # Defines the destination node that we want the pod to migrate
    destinationNode: string 
    # Service name to be changed after migration for preserving network access
    serviceName: string
status:
    # Pending | Checkpointing | Transferring | Restoring | Cleaning | Completed | Error
    migrationStatus: string
    # Message of the migration if there is any error
    migrationMessage: string
    # Name of the checkpointing file
    checkpointFile: string
    # On stages Checkpointing | Transferring | Restoring pod will be unavailable
    # For other migration types using page-server and other advanced methods this
    # period can be shortened.
    podAccessible: bool
```

## Stages
### 1. Initiation
The controller receives the newly created LPM object. By default the `migrationStatus` field is empty. Controller checks if the cluster satisfies the conditions for the migration. These conditions are given below and they can be extended.
* Does cluster has the *livepodmigration* namespace?
* Are migrators running?
* Is the pod running?
* Does the cluster has the feature gate enabled? (TODO)
* Are there any other migrations happening for the same pod? (TODO)
* Are service types supported? (TODO)
* Did the checkpoint folders are mounted to `migrators`? (TODO)

After the controller checks these conditions it will either changes the status of the migration to `Error` or `Pending`.

### 2. Triggering the Migration
At the same time migrator, and deamonsets listen to any updates on LPM. If they get an update with the status `Pending`, the migrator on the destination node will start listening for a *gRPC* connection. The one with the same node as the pod will try to initiate a *gRPC* connection. When they establish a connection the *client* migrator will change the LVM status from `Pending` to `Checkpointing`.

### 3. Checkpointing
Then client migrator stops the pod and uses [kubernetes checkpoint api](https://kubernetes.io/docs/reference/node/kubelet-checkpoint-api/) to create a checkpoint to folder `/var/lib/kubelet/checkpoint`. The host should also be mounted to the migrator so it can access it. 

After checkpointing is finnished, the client migrator changes the status from `Checkpointing` to `Transferring`.

### 4. Transferring
The client sends the checkpoint file to the server migrator. When the transfer completes server checks the file and if everything ok sends a confirmation to the client migrator. Then changes the status from `Transferring` to `Restoring`. 

### 5. Validating
The server migrator opens the checkpointing file. If any errors occur server sends a negative message and the transfer will be tried 2 more times.

### 6. Restoring
If there are no errors the server migrator will connect to the containerd socket to revive the checkpointed containers.
> Here we may run into some problems because of trying to bypass the kubelet. So for further implementations, we may want to add an extension to kubelet instead.

Here the server migrator changes the status from `Restoring` to `Cleaning`.

### 7. Cleanup
If the pod and the containers are restored then the cleanup process starts. In this stage, the server migrator sends a successful close signal to the client migrator. They clean the checkpoint and if they created other temp files that they created. And server migrator changes the status from `Cleaning` to `Completed`.