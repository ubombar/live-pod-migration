package v1alpha1

import (
	"fmt"
	"net"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
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
