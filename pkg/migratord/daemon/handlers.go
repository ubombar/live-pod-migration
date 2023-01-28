package daemon

import (
	"github.com/sirupsen/logrus"
)

func (d *daemon) incomingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusIncoming, PreparingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	//
	if *role == MigrationRoleClient {

	}

	//
	if *role == MigrationRoleServer {

	}

	return d.GetSyncer().FinishJob(migrationid, *role)
}

func (d *daemon) preparingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusPreparing, CheckpointingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

func (d *daemon) checkpointingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusCheckpointing, TransferingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

func (d *daemon) transferingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusTransfering, RestoringQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

func (d *daemon) restoringCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusRestoring, DoneQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

	return d.GetSyncer().FinishJob(migrationid, *role)
}

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

	logrus.Infoln("Migration", job.MigrationId, "finnished with error:", job.Error.Error())

	return d.GetSyncer().FinishJob(migrationid, *role)
}
