package docker

import (
	"bufio"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"io"
	"io/ioutil"
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

func (c *Controller) EnsureImage(image string) (err error){
	var reader io.ReadCloser

	output.Tracef("Running ImagePull for image %s", image)
	reader, err = c.cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if reader != nil {
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output.Trace(scanner.Text())
		}
	}

	return
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

func (c *Controller) ContainerFiles(id string, filespec string) (files []string) {
	cmd := []string{"find", "/" , "-iname", filespec}
	response, err := c.cli.ContainerExecCreate(context.Background(), id, types.ExecConfig{
		AttachStderr: false,
		AttachStdout: true,
		Cmd:          cmd,
	})

	checkErr(err)
	r, err := c.cli.ContainerExecAttach(context.Background(), response.ID, types.ExecStartCheck{})
	checkErr(err)
	defer r.Close()

	// read the output
	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		// StdCopy demultiplexes the stream into two buffers
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, r.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		checkErr(err)
		break
	case <-context.Background().Done():
		checkErr(context.Background().Err())
		break
	}
	stdout, err := ioutil.ReadAll(&outBuf)
	checkErr(err)

	return strings.Split(string(stdout), "\n")
}

func (c *Controller) ContainerExists(id string) (exists bool) {
	filters := filters.NewArgs()
	filters.Add(
		"id", id,
	)
	resp, err := c.cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: filters})
	checkErr(err)
	if len(resp) > 0 {
		output.Struct(resp)
		containerStatus := strings.Split(resp[0].Status, " ")
		status := containerStatus[0]
		output.Struct(status)
		exists = true
	}

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
