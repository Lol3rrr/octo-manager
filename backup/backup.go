package backup

import (
	"octo-manager/remote"

	"github.com/sirupsen/logrus"
)

func backup(serverDir string, remoteCon remote.Session, storageInterface storage) error {
	if len(serverDir) == 0 {
		serverDir = "."
	}
	if serverDir[len(serverDir)-1] == '/' {
		serverDir = serverDir[:len(serverDir)-1]
	}

	logrus.Infof("[Backup] Getting a list of all Files... \n")
	files, err := remoteCon.GetFiles(serverDir)
	if err != nil {
		return err
	}

	for i := range files {
		files[i].SanitizePath(serverDir)
	}

	return storageInterface.Save(files)
}
