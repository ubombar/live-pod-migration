package daemon

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (d *daemon) incomingCallback(migrationid string) error {
	_, _, _, err := FromMigrationId(migrationid, d)

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	err = d.GetRPC().NotifyPeerAboutStateChange(migrationid, StatusIncoming, StatusPreparing, nil)

	if err != nil {
		return err
	}

	d.GetQueue(PreparingQueue).Push(migrationid)

	// d.GetSyncer().RegisterJob(migrationid, StatusIncoming, PreparingQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// // CLIENT
	// if *role == MigrationRoleClient {
	// 	ins, err := d.GetDefaultContainerClient().InspectContainer(job.ClientContainerID)

	// 	if err != nil || !ins.State.Running {
	// 		logrus.Error("Container check error")
	// 		return d.GetSyncer().FinishJobWithError(migrationid, errors.New("container is not running"), *role)
	// 	}
	// }

	// // SERVER
	// // if *role == MigrationRoleServer {
	// // 	_, err = d.GetDefaultContainerClient().InspectContainer(job.ClientContainerID)

	// // 	if err == nil {
	// // 		logrus.Error("Container check error")
	// // 		return d.GetSyncer().FinishJobWithError(migrationid, errors.New("container exsists in server node"), *role)
	// // 	}
	// // }

	// logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

func (d *daemon) preparingCallback(migrationid string) error {
	// d.GetSyncer().RegisterJob(migrationid, StatusPreparing, CheckpointingQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

func (d *daemon) checkpointingCallback(migrationid string) error {
	// d.GetSyncer().RegisterJob(migrationid, StatusCheckpointing, TransferingQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// if *role == MigrationRoleClient {
	// 	checkpointPath := fmt.Sprint("/tmp/", job.MigrationId, ".tar.gz")
	// 	err := d.GetDefaultContainerClient().CheckpointContainer(job.ClientContainerID, checkpointPath)

	// 	if err != nil {
	// 		return d.GetSyncer().FinishJobWithError(migrationid, err, *role)
	// 	}
	// }

	// logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

func (d *daemon) transferingCallback(migrationid string) error {
	// d.GetSyncer().RegisterJob(migrationid, StatusTransfering, RestoringQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// if *role == MigrationRoleClient {
	// 	checkpointPath := fmt.Sprint("/tmp/", job.MigrationId, ".tar.gz")
	// 	err = d.GetDefaultContainerClient().TransferCheckpoint(checkpointPath, job.ServerIP, job.Username)

	// 	if err != nil {
	// 		return d.GetSyncer().FinishJobWithError(migrationid, err, *role)
	// 	}
	// }

	// logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

func (d *daemon) restoringCallback(migrationid string) error {
	// d.GetSyncer().RegisterJob(migrationid, StatusRestoring, DoneQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// if *role == MigrationRoleServer {
	// 	checkpointPath := fmt.Sprint("/tmp/", job.MigrationId, ".tar.gz")
	// 	_, err := d.GetDefaultContainerClient().RestoreContainer(checkpointPath, true)

	// 	if err != nil {
	// 		return d.GetSyncer().FinishJobWithError(migrationid, err, *role)
	// 	}
	// }

	// logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

func (d *daemon) doneCallback(migrationid string) error {
	// d.GetSyncer().RegisterJob(migrationid, StatusDone, NullQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// checkpointPath := fmt.Sprint("/tmp/", job.MigrationId, ".tar.gz")
	// os.Remove(checkpointPath)

	// logrus.Infoln("Migration", job.MigrationId, "finnished successfully")

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

// Prints the info and deletes all traces of the migration job since it is inactive.
func (d *daemon) errorCallback(migrationid string) error {
	// job, role, sync, err := FromMigrationId(migrationid, d)

	// d.GetSyncer().RegisterJob(migrationid, StatusDone, NullQueue)
	// job, role, err := d.getMigrationObjects(migrationid)

	// if err != nil {
	// 	return err
	// }

	// logrus.Infoln("Migration", job.MigrationId, "finnished with error:", job.Error.Error())

	// return d.GetSyncer().FinishJob(migrationid, *role)
	return nil
}

// migrations waits here unitl peer node gives it a pass
func (d *daemon) syncCallback(migrationid string) error {
	_, _, sync, err := FromMigrationId(migrationid, d)

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	if !sync.StateCheckpointReached() {
		// Add it to the queue again and try later
		d.GetQueue(SyncQueue).Push(migrationid)
		time.Sleep(100 * time.Millisecond)
		return nil
	}

	d.GetQueue(sync.GetNextState().ToQueueName()).Push(migrationid)
	return nil
}
