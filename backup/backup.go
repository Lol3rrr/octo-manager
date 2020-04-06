package backup

import (
	"octo-manager/backup/remote"

	ssh "github.com/helloyi/go-sshclient"
	"github.com/sirupsen/logrus"
)

func backup(serverDir string, sshClient *ssh.Client, storageInterface storage) error {
	if len(serverDir) == 0 {
		serverDir = "."
	}
	if serverDir[len(serverDir)-1] == '/' {
		serverDir = serverDir[:len(serverDir)-1]
	}

	logrus.Infof("[Backup] Getting a list of all Files... \n")
	files, err := remote.GetFiles(serverDir, sshClient)
	if err != nil {
		return err
	}

	for i := range files {
		files[i].SanitizePath(serverDir)
	}

	return storageInterface.Save(files)
}
