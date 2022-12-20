package v1alpha1

import (
	"context"
	"fmt"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/tmc/scp"
	pb "github.com/ubombar/live-pod-migration/pkg/migrator"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TransferCheckpointFile(job *MigrationJob, m *Migratord) error {
	key, err := ssh.ParsePrivateKey([]byte(job.PrivateKey))

	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: job.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	stringAndPort := fmt.Sprint(job.ServerIP, ":22")
	client, err := ssh.Dial("tcp", stringAndPort, config)

	if err != nil {
		return err
	}

	session, err := client.NewSession()
	defer client.Close()

	if err != nil {
		return err
	}

	migrationFolder := fmt.Sprintf(m.checkpointDir, "/", job.MigrationId)
	err = scp.CopyPath(migrationFolder, migrationFolder, session)

	// fmt.Printf("session: %v\n", session)
	// fmt.Printf("migrationFolder: %v\n", migrationFolder)

	if err != nil {
		return err
	}

	return nil
}

func restoreContainer(m *Migratord, job *MigrationJob) {
	// Check if we have the container beforehand
	err := m.Client.ContainerStart(context.Background(), job.ContainerID, types.ContainerStartOptions{
		CheckpointID:  job.MigrationId,
		CheckpointDir: m.checkpointDir,
	})

	if err != nil {
		logrus.Error(err)
		return
	}

	migratorString := fmt.Sprintf("%v:%d", job.ServerIP, job.ServerPort)
	conn, err := grpc.Dial(migratorString, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Cannot reach the migratord client on %s\n", migratorString)
		return
	}

	client := pb.NewMigratorServiceClient(conn)

	updateAndSyncJobStatus(client, m.MigrationMap, job, Done, true)
}
