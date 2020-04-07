package backup

import (
	"errors"
	"octo-manager/backup/local"
	"octo-manager/jobs"
	"octo-manager/remote"
)

func restoreLocally(stage *jobs.Stage, env jobs.Environment, remoteCon remote.Session) error {
	localDir, found := stage.GetVariable("localDir", env)
	if !found {
		return errors.New("Missing Variable: 'localDir'")
	}
	serverDir, found := stage.GetVariable("serverDir", env)
	if !found {
		return errors.New("Missing Variable: 'serverDir'")
	}

	local := &local.Storage{
		LocalDir: localDir,
	}

	return restore(serverDir, remoteCon, local)
}
