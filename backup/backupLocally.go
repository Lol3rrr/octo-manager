package backup

import (
	"errors"
	"octo-manager/backup/local"
	"octo-manager/jobs"
	"octo-manager/remote"
)

func backupLocally(stage *jobs.Stage, env jobs.Environment, remoteCon remote.Session) error {
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

	return backup(serverDir, remoteCon, local)
}
