package docker

import (
	"errors"
	"fmt"
	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
)

// RunStage is used to run a single Stage using this Module
func (m *Module) RunStage(stage *jobs.Stage, sshClient *ssh.Client, env jobs.Environment) error {
	if stage.Action == "pull" {
		imageName, found := stage.GetVariable("dockerImage", env)
		if !found {
			return errors.New("Missing Variable: 'dockerImage'")
		}

		return pullImage(sshClient, imageName)
	}
	if stage.Action == "compose up" {
		composeDir, found := stage.GetVariable("composeDir", env)
		if !found {
			return errors.New("Missing Variable: 'composeDir'")
		}

		return composeUp(sshClient, composeDir)
	}

	return fmt.Errorf("Could not find Action in 'docker'-Category: '%s'", stage.Action)

	return nil
}
