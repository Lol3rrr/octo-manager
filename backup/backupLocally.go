package backup

import (
	ssh "github.com/helloyi/go-sshclient"
	"github.com/sirupsen/logrus"
)

func backupLocally(serverDir, localDir string, sshClient *ssh.Client) error {
	logrus.Infof("[Backup][Local] Getting a list of all Files... \n")
	files, err := getFiles(serverDir, sshClient)
	if err != nil {
		return err
	}

	logrus.Infof("[Backup][Local] Saving Files locally... \n")
	err = saveFilesLocally(localDir, files)
	return err
}
