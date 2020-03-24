package backup

import (
	ssh "github.com/helloyi/go-sshclient"
	"github.com/sirupsen/logrus"
)

func backupLocally(serverDir, localDir string, sshClient *ssh.Client) error {
	if len(serverDir) == 0 {
		serverDir = "."
	}
	if serverDir[len(serverDir)-1] == '/' {
		serverDir = serverDir[:len(serverDir)-1]
	}

	logrus.Infof("[Backup][Local] Getting a list of all Files... \n")
	files, err := getFiles(serverDir, sshClient)
	if err != nil {
		return err
	}

	logrus.Infof("[Backup][Local] Saving Files locally... \n")
	return saveFilesLocally(serverDir, localDir, files)
}
