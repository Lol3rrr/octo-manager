package docker

import (
	"errors"
	"fmt"
	"octo-manager/jobs"
)

// RunStage is used to run a single Stage using this Module
func (m *Module) RunStage(ctx *jobs.Ctx) error {
	stage := ctx.Stage
	remoteSession := ctx.RemoteCon
	env := ctx.Env

	if stage.Action == "pull" {
		imageName, found := stage.GetVariable("dockerImage", env)
		if !found {
			return errors.New("Missing Variable: 'dockerImage'")
		}

		return pullImage(remoteSession, imageName)
	}
	if stage.Action == "compose up" {
		composeDir, found := stage.GetVariable("composeDir", env)
		if !found {
			return errors.New("Missing Variable: 'composeDir'")
		}

		return composeUp(remoteSession, composeDir)
	}

	return fmt.Errorf("Could not find Action in 'docker'-Category: '%s'", stage.Action)
}
