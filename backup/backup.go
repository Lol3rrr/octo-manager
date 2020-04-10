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

	logrus.Infof("[Backup] Starting Backup of '%s' \n", serverDir)
	logrus.Infof("[Backup] Getting Files... \n")
	files, err := remoteCon.GetFiles(serverDir)
	if err != nil {
		return err
	}

	logrus.Infof("[Backup] Sanitizing Paths... \n")
	for i := range files {
		files[i].SanitizePath(serverDir)
	}

	logrus.Infof("[Backup] Saving Files... \n")
	return storageInterface.Save(files)
}
