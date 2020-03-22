package docker

import (
	ssh "github.com/helloyi/go-sshclient"
)

func ComposeUp(sshClient *ssh.Client, dir string) error {
	cmdString := "docker-compose up -d;"
	if len(dir) > 0 {
		cmdString = "cd " + dir + ";" + cmdString
	}

	err := sshClient.Cmd(cmdString).Run()
	if err != nil {
		return err
	}

	return nil
}
