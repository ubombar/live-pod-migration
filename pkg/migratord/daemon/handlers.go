package daemon

import (
	"github.com/sirupsen/logrus"
)

// This handle check the following conditions to allow the migration to happen.
// * There should be a container with the following container id
// * There should not be any active migration with the same client container
func (d *daemon) incomingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusIncoming, PreparingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	// Client
	if *role == MigrationRoleClient {
		// // Get the container with the container id
		// _, err = d.GetDefaultContainerClient().GetContainerInfo(job.ClientContainerID)

		// if err != nil {
		// 	return d.GetSyncer().FinishJobWithError(migrationid, err, *role)
		// }

	}

	return d.GetSyncer().FinishJob(migrationid, *role)
}

// Handle migration
func (d *daemon) preparingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusPreparing, CheckpointingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

// Handle migration
func (d *daemon) checkpointingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusCheckpointing, TransferingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

// Handle migration
func (d *daemon) transferingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusTransfering, RestoringQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

// Handle migration
func (d *daemon) restoringCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusRestoring, DoneQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

// Handle migration
func (d *daemon) doneCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusDone, NullQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished successfully")

	return d.GetSyncer().FinishJob(migrationid, *role)
}

// Prints the info and deletes all traces of the migration job since it is inactive.
func (d *daemon) errorCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusDone, NullQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished with error")

	return d.GetSyncer().FinishJob(migrationid, *role)
}
