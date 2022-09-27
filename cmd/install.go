// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"gopkg.in/yaml.v2"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
)

type Controller struct {
	cli *client.Client
}

func NewController() (c *Controller, err error) {
	c = new(Controller)

	c.cli, err = client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Controller) EnsureImage(image string) (err error) {
	reader, err := c.cli.ImagePull(context.Background(), image, types.ImagePullOptions{})

	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}

func (c *Controller) ContainerRun(image string, env[] string, command []string) (id string, err error) {

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"1433/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "1433",
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
		return "", err
	}

	return resp.ID, nil
}

type InstallArguments struct {
	Name string
}

var installArguments InstallArguments

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install SQL Server and Tools",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var imageName =  "mcr.microsoft.com/mssql/server:2022-latest"
		// Generate 100 character sa Password
		rand.Seed(time.Now().Unix())
		password := generatePassword(100, 2, 2, 2)

		env := []string{"ACCEPT_EULA=Y", fmt.Sprintf("SA_PASSWORD=%s", password)}

		c, err := NewController()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Installing SQL Server ('2022-latest')\n")

		err = c.EnsureImage(imageName)
		if err != nil {
			fmt.Println(err)
		}

		 id, err := c.ContainerRun(imageName, env, []string{})
		if err != nil {
			fmt.Println(err)
		}

		updateConfig(id, password)

		fmt.Printf("SQL Server installed (id: '%s', current context: 'sa@sql1')\n", id[len(id)-12:])
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func updateConfig(id string, password string) {
	var config Sqlconfig

	config.ApiVersion = "v1"
	config.Kind = "Config"
	config.CurrentContext = "sa@sql1"

	config.Endpoints = []Endpoint{
		{Name: "sql1",
			DockerDetails: DockerDetails{
				ContainerId: id},
			EndpointDetails: EndpointDetails{
				Address: "localhost",
				Port:    1433}}}

	config.Contexts = []Context{
		{Name: "sa@sql1",
			ContextDetails: ContextDetails{
				Endpoint: "localhost",
				User:     "sa"}}}

	config.Users = []User{
		{Name: "sa", Password: base64.StdEncoding.EncodeToString([]byte(password)),
			UserDetails: UserDetails{
				Username: "sa",
				Password: base64.StdEncoding.EncodeToString([]byte(password))}}}

	b, err := yaml.Marshal(&config)

	// BUGBUG: Should be able to get tihs from viper, viper.ConfigFileUsed() is
	// returning empty
	var configFile = os.Getenv("USERPROFILE")
	configFile = filepath.Join(configFile, ".sqlcmd", "sqlconfig")

	// File has to exist for WriteConfig to work
	if !fileExists(configFile) {
		mkdirDir(filepath.Join(os.Getenv("USERPROFILE"), ".sqlcmd"))
		f, err := os.Create(configFile)
		if err != nil {
			fmt.Println(err)
		}
		f.Close()
	}

	viper.ReadConfig(bytes.NewReader(b))

	err = viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func fileExists(filename string) (exists bool) {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func mkdirDir(folder string) {
	if folder == "" {
		return
	}
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Unable to create folder '%v'. %v", folder, err))
		}
	}
}

// https://golangbyexample.com/generate-random-password-golang/
var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
