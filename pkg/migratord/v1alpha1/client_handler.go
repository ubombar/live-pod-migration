package v1alpha1

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	pb "github.com/ubombar/live-pod-migration/pkg/migrator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientMigationHandler struct {
	parent *Migratord
}

// Creates different adapters depending on the migration job
func (h *clientMigationHandler) PerformMigration(migration *MigrationJob) {
	switch migration.Method {
	case Basic:
		handleBasicMigrationJob(h.parent, migration)
	case Postcopy:
		// Not implemented yet
		handleUnknownMigrationJob(h.parent, migration)
	case Precopy:
		// Not implemented yet
		handleUnknownMigrationJob(h.parent, migration)
	default:
		// Error
		handleUnknownMigrationJob(h.parent, migration)
	}
}

func updateAndSyncJobStatus(client pb.MigratorServiceClient, mmap *MigrationMap, job *MigrationJob, newStatus MigrationStatus, running bool) (*MigrationJob, error) {
	// Change the status in server
	_, err := client.UpdateMigrationStatus(context.Background(), &pb.UpdateMigrationStatusRequest{
		MigrationId: job.MigrationId,
		NewStatus:   string(newStatus),
		NewRunning:  running,
	})

	if err != nil {
		return job, err
	}

	// Change the status in client
	newJobs := *job
	newJobs.Status = newStatus
	newJobs.Running = running
	mmap.Save(&newJobs)

	return &newJobs, nil
}

// Handles unknown or not implemented migration job
func handleUnknownMigrationJob(m *Migratord, job *MigrationJob) {
	logrus.Warn("Migration job method is not implemented.")
}

// Do the basic migration job. This consists of Checkpointing, Transfering and Restoring.
func handleBasicMigrationJob(m *Migratord, job *MigrationJob) {
	migratorString := fmt.Sprintf("%v:%d", job.ServerIP, job.ServerPort)
	conn, err := grpc.Dial(migratorString, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Cannot reach the migratord client on %s\n", migratorString)
		return
	}

	client := pb.NewMigratorServiceClient(conn)

	// Checkpointing state
	job, err = updateAndSyncJobStatus(client, m.MigrationMap, job, Checkpointing, false)
	if err != nil {
		logrus.Error(err)
	}

	// Create checkpoint
	err = m.Client.CheckpointCreate(context.Background(), job.ContainerID, types.CheckpointCreateOptions{
		CheckpointDir: m.checkpointDir,
		CheckpointID:  job.MigrationId,
		Exit:          true,
	})

	if err != nil {
		logrus.Error(err)
		return
	}

	// Transfering state
	logrus.Infoln("Migration", job.MigrationId, "finnished checkpointing now transferring.")
	job, err = updateAndSyncJobStatus(client, m.MigrationMap, job, Transfering, false)
	if err != nil {
		logrus.Error(err)
	}

	// Restoring state
	logrus.Infoln("Migration", job.MigrationId, "finnished transfering now restoring.")
	job, err = updateAndSyncJobStatus(client, m.MigrationMap, job, Restoring, false)
	if err != nil {
		logrus.Error(err)
	}

	// Done state
	logrus.Infoln("Migration", job.MigrationId, "finnished restoring.")
	job, err = updateAndSyncJobStatus(client, m.MigrationMap, job, Done, false)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infoln("Migration", job.MigrationId, "complete!")
}
