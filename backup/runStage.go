package backup

import (
	"errors"
	"fmt"
	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
)

// RunStage is used to run a single Stage using the Backup-Module
func (m *Module) RunStage(stage *jobs.Stage, sshClient *ssh.Client, env jobs.Environment) error {
	if stage.Action == "local" {
		serverDir, found := stage.GetVariable("serverDir", env)
		if !found {
			return errors.New("Missing Variable: 'serverDir'")
		}
		localDir, found := stage.GetVariable("localDir", env)
		if !found {
			return errors.New("Missing Variable: 'localDir'")
		}

		return backupLocally(serverDir, localDir, sshClient)
	}

	return fmt.Errorf("Could not find Action in 'backup'-Category: '%s'", stage.Action)
}
