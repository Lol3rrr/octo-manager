package docker

import (
	"errors"

	ssh "github.com/helloyi/go-sshclient"
)

func PullImage(sshClient *ssh.Client, image string) error {
	if len(image) <= 0 {
		return errors.New("Image Parameter can not be empty")
	}

	cmdString := "docker pull " + image
	_, err := sshClient.Cmd(cmdString).Output()
	if err != nil {
		return err
	}

	return nil
}
