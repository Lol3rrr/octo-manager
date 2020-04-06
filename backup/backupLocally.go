package backup

import (
	"errors"
	"octo-manager/backup/local"
	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
)

func backupLocally(stage *jobs.Stage, env jobs.Environment, sshClient *ssh.Client) error {
	serverDir, found := stage.GetVariable("serverDir", env)
	if !found {
		return errors.New("Missing Variable: 'serverDir'")
	}
	localDir, found := stage.GetVariable("localDir", env)
	if !found {
		return errors.New("Missing Variable: 'localDir'")
	}

	local := &local.Storage{
		LocalDir: localDir,
	}

	return backup(serverDir, sshClient, local)
}
