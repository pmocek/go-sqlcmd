package docker

import (
	"bufio"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"strconv"
	"strings"
)

type Controller struct {
	cli *client.Client
}

func NewController() (c *Controller) {
	var err error
	c = new(Controller)
	c.cli, err = client.NewClientWithOpts(client.FromEnv)
	checkErr(err)

	return
}

func (c *Controller) EnsureImage(image string) {
	reader, err := c.cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	checkErr(err)
	defer reader.Close()

	//io.Copy(os.Stdout, reader)
}

func (c *Controller) ContainerRun(image string, env[] string, port int, command []string) (id string, err error) {
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("1433/tcp"): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: strconv.Itoa(port),
				},
			},
		},
	}

	resp, err := c.cli.ContainerCreate(context.Background(), &container.Config{
		Tty:   true,
		Image: image,
		Cmd:   command,
		Env:   env,
	}, hostConfig, nil, nil, "")

	if err != nil {
		return "", err
	}

	err = c.cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return resp.ID, err
	}

	return resp.ID, nil
}

// ContainerWaitForLogEntry waits for text substring in containers logs
func (c *Controller) ContainerWaitForLogEntry(id string, text string) {
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: false,
		Since:      "",
		Until:      "",
		Timestamps: false,
		Follow:     true,
		Tail:       "",
		Details:    false,
	}

	// Wait for server to start up
	reader, err := c.cli.ContainerLogs(context.Background(), id, options)
	checkErr(err)
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		output.Tracef("ERRORLOG: " + scanner.Text())
		if strings.Contains(scanner.Text(), text) {
			break
		}
	}
}

func (c *Controller) ContainerStop(id string) (err error) {
	err = c.cli.ContainerStop(context.Background(), id, nil)
	return
}

func (c *Controller) ContainerRemove(id string) (err error) {
	options := types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks:   false,
		Force:         false,
	}

	err = c.cli.ContainerRemove(context.Background(),id, options)

	return
}
