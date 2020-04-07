package backup

import (
	"octo-manager/backup/remote"
	sshRemote "octo-manager/remote"

	"github.com/sirupsen/logrus"
)

func backup(serverDir string, remoteCon sshRemote.Session, storageInterface storage) error {
	if len(serverDir) == 0 {
		serverDir = "."
	}
	if serverDir[len(serverDir)-1] == '/' {
		serverDir = serverDir[:len(serverDir)-1]
	}

	logrus.Infof("[Backup] Getting a list of all Files... \n")
	files, err := remote.GetFiles(serverDir, remoteCon)
	if err != nil {
		return err
	}

	for i := range files {
		files[i].SanitizePath(serverDir)
	}

	return storageInterface.Save(files)
}
