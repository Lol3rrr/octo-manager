package backup

import (
	"errors"
	"octo-manager/backup/local"
	"octo-manager/jobs"
	"octo-manager/remote"
)

func deleteLocally(stage *jobs.Stage, env jobs.Environment, remoteCon remote.Session) error {
	localDir, found := stage.GetVariable("localDir", env)
	if !found {
		return errors.New("missing Variable: 'localDir'")
	}

	local := &local.Storage{
		LocalDir: localDir,
	}

	return delete(stage, env, local)
}
