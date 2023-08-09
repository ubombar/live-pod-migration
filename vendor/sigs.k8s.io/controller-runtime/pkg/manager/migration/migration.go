package migration

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

type LiveMigrationClient interface {
	Checkpoint(containerId string, checkpointdir string) error
	Transfer(checkpointdir string) error
	Restore(checkpointdir string) error
}

type liveMigrationClient struct {
	LiveMigrationClient
	client *containerd.Client
}

func NewLiveMigrationClient() (LiveMigrationClient, error) {
	client, err := containerd.New("/run/containerd/containerd.sock")

	if err != nil {
		return nil, err
	}

	return &liveMigrationClient{
		client: client,
	}, nil
}

// func (l* liveMigrationClient) getAllContainers()

func (l *liveMigrationClient) Checkpoint(containerId string, checkpointdir string) error {
	// ctx := namespaces.WithNamespace(context.Background(), "moby")

	namespacess, err := l.client.NamespaceService().List(context.Background())

	if err != nil {
		return err
	}

	for _, namespace := range namespacess {
		context := namespaces.WithNamespace(context.Background(), namespace)

		fmt.Printf("namespace: %v\n", namespace)
		containers, err := l.client.Containers(context)

		if err != nil {
			continue
		}

		for _, container := range containers {
			fmt.Printf("\t%v\n", container.ID())
		}
	}

	return nil

	// container, err := l.client.LoadContainer(ctx, containerId)

	// if err != nil {
	// 	return err
	// }

	// image, err := container.Image(ctx)

	// if err != nil {
	// 	return err
	// }

	// checkpointImage, err := container.Checkpoint(ctx, image.Name())

	// fmt.Printf("checkpointImage: %v\n", checkpointImage)
	// return nil
}

func (l *liveMigrationClient) Transfer(checkpointdir string) error {
	return nil
}

func (l *liveMigrationClient) Restore(checkpointdir string) error {
	return nil
}
