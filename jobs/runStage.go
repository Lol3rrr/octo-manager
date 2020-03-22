package jobs

import (
	"errors"
	"fmt"

	"octo-manager/docker"

	ssh "github.com/helloyi/go-sshclient"
)

func (stage *Stage) Run(sshClient *ssh.Client, env Environment) error {
	fmt.Printf("[Stage][%s] Starting... \n", stage.Name)
	defer fmt.Printf("[Stage][%s] Done \n", stage.Name)

	if stage.Category == "docker" {
		if stage.Action == "pull" {
			imageName, found := env["image"]
			if !found {
				return errors.New("Does not contain all needed Env-Variables")
			}

			return docker.PullImage(sshClient, imageName)
		}
		if stage.Action == "compose up" {
			composeDir, found := stage.Variables["dir"]
			if !found {
				return errors.New("Does not contain all needed Variables")
			}

			return docker.ComposeUp(sshClient, composeDir)
		}

		return errors.New("Could not find Action in 'docker'-Category")
	}

	return errors.New("Could not find Category")
}
