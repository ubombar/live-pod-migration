# Steps of the Migration
This project contains two binary files that will be deployed to the cluster. First one is the `custom controller` and the second one is the `migrator`. The CRD for managing this operation is named `LivePodMigration`. Here is the inside of the CRD:

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
    # Pending | Checkpointing | Transferring | Restoring | Completed | Error
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
Controller receives the newly created LPM object. By default the `migrationStatus` field is empty. Controller checks if the cluster satisfys the conditions for the migration. These conditions are given below and they can be extended.
* Does cluster has the *livepodmigration* namespace?
* Are migrators running?
* Is pod running?
* Does cluster has the featuregate enabled? (TODO)
* Are there any other migrations happening for the same pod? (TODO)
* Are service types supported? (TODO)
* Did the checkpoint folders are mounted to `migrators`? (TODO)

After controller checking these conditions it will either changes the status of the migration to `Error` or `Pending`.

### 2. Migration Trigger
At the same time migrator deamonsets listen any updates on LVM. If they get an update with the status `Pending`, the migrator on the destination node will start listening for a *gRPC* connection. The one with the same node as the pod will try to initiate a *gRPC* connection. When they establish a connection the *client* migrator will change the LVM status from `Pending` to `Transferring`.