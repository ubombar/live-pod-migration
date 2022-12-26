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
func (c *dockerClient) GetContainerInfo(cid string) (*ContainerInfo, error) {
	con, err := c.cli.ContainerInspect(context.Background(), cid)

	if err != nil {
		return nil, err
	}

	var containerId string
	var imageid string

	if strings.Contains(con.ID, ":") {
		containerId = strings.Split(con.ID, ":")[1]
	}

	if strings.Contains(con.Image, ":") {
		imageid = strings.Split(con.Image, ":")[1]
	}

	info := &ContainerInfo{
		ContainerId:    containerId,
		ImageId:        imageid,
		ContainerNames: []string{con.Name},
	}

	return info, nil
}

// Get checkpoint info
func (c *dockerClient) GetCheckpointInfo(checkpointid string) (*CheckpointInfo, error) {
	return nil, nil
}
