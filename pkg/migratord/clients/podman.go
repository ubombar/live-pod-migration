package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

type podmanClient struct {
	Client
	containerSocket string
}

func NewPodmanClient() *podmanClient {
	c := &podmanClient{
		containerSocket: "/run/podman/podman.sock",
	}

	if _, err := os.Stat(c.containerSocket); os.IsNotExist(err) {
		return nil
	}

	return c
}

// Get runtime eg 'podman'
func (c *podmanClient) Runtime() string {
	return ClientPodman
}

// Get version of the runtime
func (c *podmanClient) Version() string {
	cmd := exec.Command("podman", "--version")
	stdout, err := cmd.Output()

	if err != nil {
		return "unknown"
	}

	splitted := strings.Split(string(stdout), " ")

	if len(splitted) != 3 {
		return "unknown"
	}

	return splitted[2][:len(splitted[2])-1]
}

func (c *podmanClient) InspectContainer(containerId string) (*ContainerInspectResult, error) {
	cmd := exec.Command("podman", "container", "inspect", containerId)
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	containerObject := []ContainerInspectResult{}
	err = json.Unmarshal(output, &containerObject)

	if err != nil {
		return nil, errors.New("container does not exists")
	}

	return &containerObject[0], nil
}

func (c *podmanClient) CheckpointContainer(containerId string, checkpointPath string) error {
	cmd := exec.Command("podman", "container", "checkpoint", "--export", checkpointPath, containerId)
	output, err := cmd.Output()

	if err != nil {
		return err
	}

	outputString := string(output)

	if strings.Contains(outputString, "Error: ") {
		return errors.New(outputString)
	}

	return nil
}

func (c *podmanClient) RestoreContainer(checkpointPath string, randomizeName bool) (*ContainerInspectResult, error) {
	var cmd *exec.Cmd
	randomName := randomString(16)

	if randomizeName {
		cmd = exec.Command("podman", "container", "restore", "--import", checkpointPath, "--name", randomName)
	} else {
		cmd = exec.Command("podman", "container", "restore", "--import", checkpointPath)
	}
	output, err := cmd.Output()

	if err != nil {
		return nil, errors.New("cannot connect to podman")
	}

	outputString := string(output)

	if strings.Contains(outputString, "Error: ") {
		return nil, errors.New(outputString)
	}

	result, err := c.InspectContainer(randomName)

	return result, err
}

func (c *podmanClient) ClearContainer(containerId string) {

}

func (c *podmanClient) TransferCheckpoint(checkpointPath string, serverAddress string) error {
	cmd := exec.Command("scp", checkpointPath, fmt.Sprint(serverAddress, ":", checkpointPath))

	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func randomString(n int) string {
	var letters = []rune("0123456789abcdef")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
