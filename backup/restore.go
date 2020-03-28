package backup

import (
	"octo-manager/backup/remote"

	ssh "github.com/helloyi/go-sshclient"
)

func restore(serverDir string, sshClient *ssh.Client, storageInterface storage) error {
	files, err := storageInterface.LoadLatestFiles()
	if err != nil {
		return err
	}

	return remote.WriteFiles(files, serverDir, sshClient)
}
