package backup

import (
	"fmt"
	"octo-manager/jobs"
)

// RunStage is used to run a single Stage using the Backup-Module
func (m *Module) RunStage(ctx *jobs.Ctx) error {
	stage := ctx.Stage
	sshClient := ctx.SSHClient
	env := ctx.Env

	if stage.Action == "save-local" {
		return backupLocally(stage, env, sshClient)
	}
	if stage.Action == "save-googleDrive" {
		return backupGoogleDrive(stage, env, sshClient)
	}

	if stage.Action == "restore-local" {
		return restoreLocally(stage, env, sshClient)
	}
	if stage.Action == "restore-googleDrive" {
		return restoreGoogleDrive(stage, env, sshClient)
	}

	return fmt.Errorf("Could not find Action in 'backup'-Category: '%s'", stage.Action)
}
