package backup

import (
	"octo-manager/backup/remote"
	sshRemote "octo-manager/remote"
)

func restore(serverDir string, remoteCon sshRemote.Session, storageInterface storage) error {
	files, err := storageInterface.LoadLatestFiles()
	if err != nil {
		return err
	}

	return remote.WriteFiles(files, serverDir, remoteCon)
}
