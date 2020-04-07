package docker

import (
	"errors"
	"octo-manager/remote"
)

func pullImage(remoteCon remote.Session, image string) error {
	if len(image) <= 0 {
		return errors.New("image Parameter can not be empty")
	}

	cmdString := "docker pull " + image
	_, err := remoteCon.Command(cmdString)
	if err != nil {
		return err
	}

	return nil
}
