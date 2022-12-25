package daemon

import "github.com/sirupsen/logrus"

// Handle migration
func (d *daemon) incomingCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusIncoming, PreparingQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	// Check for any problems

	logrus.Infoln("Migration", job.MigrationId, "finnished stage", job.Status)

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

// Handle migration
func (d *daemon) errorCallback(migrationid string) error {
	d.GetSyncer().RegisterJob(migrationid, StatusDone, NullQueue)
	job, role, err := d.getMigrationObjects(migrationid)

	if err != nil {
		return err
	}

	logrus.Infoln("Migration", job.MigrationId, "finnished successfully")

	return d.GetSyncer().FinishJob(migrationid, *role)
}
