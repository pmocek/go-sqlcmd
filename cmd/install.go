// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"bufio"
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
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/billgraziano/dpapi"

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

func (c *Controller) ContainerWaitForLogEntry(id string, text string) (err error) {
	// Wait for "SQL Server is now ready for client connections"
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
	cobra.CheckErr(err)
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			break
		}
	}

	return
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
		var config Sqlconfig

		viper.Unmarshal(&config)

		port := findFreePortForTds(config)

		var imageName =  "mcr.microsoft.com/mssql/server:2022-latest"
		// Generate 100 character sa Password
		rand.Seed(time.Now().Unix())
		password := generatePassword(100, 2, 2, 2)

		env := []string{"ACCEPT_EULA=Y", fmt.Sprintf("SA_PASSWORD=%s", password)}

		c, err := NewController()
		cobra.CheckErr(err)

		fmt.Printf("Installing SQL Server ('2022-latest')\n")

		err = c.EnsureImage(imageName)
		cobra.CheckErr(err)

		id, err := c.ContainerRun(imageName, env, port, []string{})
		if err != nil {
			// Remove the container, because we haven't persisted to config yet, so
			// uninstall won't work yet
			c.ContainerRemove(id)
		}
		cobra.CheckErr(err)

		updateConfig(id, port, password)

		err = c.ContainerWaitForLogEntry(id, "SQL Server is now ready for client connections")
		cobra.CheckErr(err)

		fmt.Printf("SQL Server installed (id: '%s', current context: '%v')\n",
			id[len(id)-12:],
			config.CurrentContext)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func updateConfig(id string, portNumber int, password string) {
	var config Sqlconfig
	err := viper.Unmarshal(&config)
	cobra.CheckErr(err)

	encryptedPassword := password
	if runtime.GOOS == "windows" {
		encryptedPassword, err= dpapi.Encrypt(password)
		cobra.CheckErr(err)
	} else {
		// TODO: MacOS (keychain) and Linux (not sure?)
	}

	endPointName := findEndpointName(config)

	config.ApiVersion = "v1"
	config.Kind = "Config"
	config.CurrentContext = "sa@" + endPointName

	config.Endpoints = append(config.Endpoints, Endpoint{
		DockerDetails:   DockerDetails{ContainerId: id},
		EndpointDetails: EndpointDetails{
			Address: "localhost",
			Port:    portNumber,
		},
		Name:            endPointName,
	})

	config.Contexts = append(config.Contexts, Context{
		ContextDetails: ContextDetails{
			Endpoint: endPointName,
			User:     "sa@" + endPointName,
		},
		Name:           "sa@" + endPointName,
	})

	config.Users = append(config.Users, User{
		UserDetails: UserDetails{
			Username: "sa",
			Password: base64.StdEncoding.EncodeToString([]byte(encryptedPassword)),
		},
		Name:        "sa@" + endPointName,
	})

	err = saveConfig(config)
	cobra.CheckErr(err)
}

func findFreePortForTds(config Sqlconfig) (portNumber int) {
	const startingPortNumber = 1433

	portNumber = startingPortNumber

	for {
		foundFreePortNumber := true
		for _, endpoint := range config.Endpoints {
			if endpoint.Port == portNumber {
				foundFreePortNumber = false
				break
			}
		}

		if foundFreePortNumber == true {
			break
		}

		portNumber++

		if portNumber == 5000 {
			panic("Did not find an available port")
		}
	}

	return
}

func findEndpointName(config Sqlconfig) (endPointName string) {
	var postfixNumber = 1

	for {
		endPointName = fmt.Sprintf("sql%v", strconv.Itoa(postfixNumber))
		foundAvailableEndpointName := true
		for _, endpoint := range config.Endpoints {
			if endpoint.Name == endPointName {
				foundAvailableEndpointName = false
				break
			}
		}

		if foundAvailableEndpointName == true {
			break
		}

		postfixNumber++

		if postfixNumber == 5000 {
			panic("Did not find an available endpoint name")
		}
	}

	return endPointName
}

func saveConfig(config Sqlconfig) (err error) {
	b, err := yaml.Marshal(&config)
	cobra.CheckErr(err)

	// BUGBUG: Should be able to get this from viper, viper.ConfigFileUsed() is
	// returning empty
	var configFile = os.Getenv("USERPROFILE")
	configFile = filepath.Join(configFile, ".sqlcmd", "sqlconfig")

	viper.ReadConfig(bytes.NewReader(b))

	// File has to exist for WriteConfig to work
	if !fileExists(configFile) {
		mkdirDir(filepath.Join(os.Getenv("USERPROFILE"), ".sqlcmd"))
		f, err := os.Create(configFile)
		defer f.Close()
		cobra.CheckErr(err)
	}

	err = viper.WriteConfig()
	cobra.CheckErr(err)

	return
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
		cobra.CheckErr(err)
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
