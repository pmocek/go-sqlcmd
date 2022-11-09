package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"testing"
)

func TestController_ListTags(t *testing.T) {
	const registry = "mcr.microsoft.com"
	const repo = "mssql/server"

	ListTags(repo, "https://" + registry)
}

func TestController_EnsureImage(t *testing.T) {
	const registry = "docker.io"
	const repo = "library/alpine"
	const tag = "latest"
	const port = 0

	imageName := fmt.Sprintf(
		"%s/%s:%s",
		registry,
		repo,
		tag)

	type fields struct {
		cli *client.Client
	}
	type args struct {
		image string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"default", fields{nil}, args{imageName}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewController()
			c.EnsureImage(tt.args.image)
			id, err := c.ContainerRun(
				tt.args.image,
				[]string{},
				port,
				[]string{"ash", "-c", "echo 'Hello World'; sleep 1"},
			)
			checkErr(err)
			c.ContainerWaitForLogEntry(id, "Hello World")
			c.ContainerExists(id)
			c.ContainerFiles(id, "*.mdf")
			c.ContainerStop(id)
			c.ContainerRemove(id)
		})
	}
}
