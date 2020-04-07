package backup

import (
	"fmt"
	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
)

// RunStage is used to run a single Stage using the Backup-Module
func (m *Module) RunStage(stage *jobs.Stage, sshClient *ssh.Client, env jobs.Environment) error {
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
