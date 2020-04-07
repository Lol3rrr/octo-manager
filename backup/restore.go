package backup

import (
	"octo-manager/remote"
)

func restore(serverDir string, remoteCon remote.Session, storageInterface storage) error {
	files, err := storageInterface.LoadLatestFiles()
	if err != nil {
		return err
	}

	return remoteCon.WriteFiles(files, serverDir)
}
