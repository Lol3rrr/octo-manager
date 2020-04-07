package remote

import (
	"octo-manager/backup/general"
	"octo-manager/remote"
	"path/filepath"
)

func WriteFiles(files []general.File, serverDir string, remoteCon remote.Session) error {
	for _, tmpFile := range files {
		tmpPath := filepath.Join(serverDir, tmpFile.Path)
		err := checkDir(filepath.Dir(tmpPath), remoteCon)
		if err != nil {
			return err
		}

		err = writeToFile(tmpFile.Content, tmpPath, remoteCon)
		if err != nil {
			return err
		}
	}

	return nil
}
