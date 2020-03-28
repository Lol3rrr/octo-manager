package remote

import (
	"octo-manager/backup/general"
	"path/filepath"

	ssh "github.com/helloyi/go-sshclient"
)

func WriteFiles(files []general.File, serverDir string, sshClient *ssh.Client) error {
	for _, tmpFile := range files {
		tmpPath := filepath.Join(serverDir, tmpFile.Path)
		err := checkDir(filepath.Dir(tmpPath), sshClient)
		if err != nil {
			return err
		}

		err = writeToFile(tmpFile.Content, tmpPath, sshClient)
		if err != nil {
			return err
		}
	}

	return nil
}
