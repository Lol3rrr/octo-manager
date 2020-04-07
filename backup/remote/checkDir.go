package remote

import (
	"fmt"
	"octo-manager/remote"
	"path/filepath"
	"strings"
)

func checkDir(dir string, remoteCon remote.Session) error {
	dirs := strings.Split(dir, "/")
	for i := range dirs {
		tmpPath := filepath.Join(dirs[:i+1]...)

		cdCmd := fmt.Sprintf("cd %s", tmpPath)
		_, err := remoteCon.Command(cdCmd)
		if err != nil {
			mkdirCmd := fmt.Sprintf("mkdir %s", tmpPath)
			_, err = remoteCon.Command(mkdirCmd)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
