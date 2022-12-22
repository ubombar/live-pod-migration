package clients

import (
	"context"
	"strings"

	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

type dockerClient struct {
	Client
	cli *client.Client
}

func NewDockerClient() *dockerClient {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		logrus.Errorln("cannot connect to docker: ", err)
		return nil
	}

	c := &dockerClient{
		cli: cli,
	}

	return c
}

// Get runtime eg 'docker'
func (c *dockerClient) Runtime() string {
	return ClientDocker
}

// Get version of the runtime
func (c *dockerClient) Version() string {
	return c.cli.ClientVersion()
}

// Get container info
func (c *dockerClient) GetContainerInfo(containerId string) (*ContainerInfo, error) {
	con, err := c.cli.ContainerInspect(context.Background(), containerId)

	if err != nil {
		return nil, err
	}

	info := &ContainerInfo{
		ContainerId:    strings.Split(con.ID, ":")[1],
		ImageId:        strings.Split(con.Image, ":")[1],
		ContainerNames: []string{con.Name},
	}

	return info, nil
}

// Get checkpoint info
func (c *dockerClient) GetCheckpointInfo(checkpointid string) (*CheckpointInfo, error) {
	return nil, nil
}
