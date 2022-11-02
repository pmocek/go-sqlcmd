package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"testing"
)

func TestController_EnsureImage(t *testing.T) {
	const registry = "mcr.microsoft.com"
	const repo = "mssql/server"
	const tag = "latest"
	const saPassword = "123456789abcde!!"

	env := []string{"ACCEPT_EULA=Y", fmt.Sprintf("SA_PASSWORD=%s", saPassword)}

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
		{ "default", fields{nil}, args{imageName}, true },
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewController()
			ListTags("azure-sql-edge", "https://mcr.microsoft.com")
			c.EnsureImage(tt.args.image)
			id, err := c.ContainerRun(tt.args.image, env, 1499, []string{})
			checkErr(err)
			c.ContainerWaitForLogEntry(id, "The default language")
			c.ContainerExists(id)
			c.ContainerFiles(id, "*.mdf")
			c.ContainerStop(id)
			c.ContainerRemove(id)
		})
	}
}
