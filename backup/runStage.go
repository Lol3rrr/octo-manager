package backup

import (
	"fmt"
	"octo-manager/jobs"
)

// RunStage is used to run a single Stage using the Backup-Module
func (m *Module) RunStage(ctx *jobs.Ctx) error {
	stage := ctx.Stage
	remoteSession := ctx.RemoteCon
	env := ctx.Env

	if stage.Action == "save-local" {
		return backupLocally(stage, env, remoteSession)
	}
	if stage.Action == "save-googleDrive" {
		return backupGoogleDrive(stage, env, remoteSession)
	}

	if stage.Action == "restore-local" {
		return restoreLocally(stage, env, remoteSession)
	}
	if stage.Action == "restore-googleDrive" {
		return restoreGoogleDrive(stage, env, remoteSession)
	}

	if stage.Action == "delete-local" {
		return deleteLocally(stage, env, remoteSession)
	}

	return fmt.Errorf("could not find Action in 'backup'-Category: '%s'", stage.Action)
}
