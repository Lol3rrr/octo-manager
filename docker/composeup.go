package docker

import (
	"octo-manager/remote"
)

func composeUp(remoteCon remote.Session, dir string) error {
	cmdString := "docker-compose up -d;"
	if len(dir) > 0 {
		cmdString = "cd " + dir + ";" + cmdString
	}

	_, err := remoteCon.Command(cmdString)
	if err != nil {
		return err
	}

	return nil
}
