package v1alpha1

import (
	"context"
	"fmt"

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
		NewStatus:   string(Checkpointing),
		NewRunning:  false,
	})

	// logrus.Println("Updating the status", job)

	if err != nil {
		logrus.Warn("Dailed to update peer status")
		return job, err
	}

	// Change the status in client
	job.Status = newStatus
	job.Running = running
	mmap.Save(job)

	return job, nil
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

	_ = pb.NewMigratorServiceClient(conn)

	fmt.Printf("%v\n", job)

	// // Change the status to Checkpointing
	// job, _ = updateAndSyncJobStatus(client, m.MigrationMap, job, Checkpointing, false)

	// // Change the status to Transfering
	// job, _ = updateAndSyncJobStatus(client, m.MigrationMap, job, Transfering, false)

	// // Change the status to Restoring
	// job, _ = updateAndSyncJobStatus(client, m.MigrationMap, job, Restoring, false)

	// // Change the status to Done
	// job, _ = updateAndSyncJobStatus(client, m.MigrationMap, job, Done, true)
}
