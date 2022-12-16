package v1alpha1

import (
	"context"
	"errors"

	"github.com/docker/docker/client"
)

// This will run as a system deamon. Checks incoming client and server migration requests.
type Migratord struct {
	Client           *client.Client
	OSType           string
	ClientAPIVersion string

	// Incoming queue

}

func NewMigratord() (*Migratord, error) {
	cl, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	ping, err := cl.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	if !ping.Experimental {
		return nil, errors.New("experimental mode is not eabled on node")
	}

	migratord := &Migratord{
		Client:           cl,
		OSType:           ping.OSType,
		ClientAPIVersion: ping.APIVersion,
	}

	return migratord, nil
}
