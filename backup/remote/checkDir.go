package remote

import (
	"fmt"
	"path/filepath"
	"strings"

	ssh "github.com/helloyi/go-sshclient"
)


func checkDir(dir string, sshClient *ssh.Client) error {
	dirs := strings.Split(dir, "/")
	for i := range dirs {
		tmpPath := filepath.Join(dirs[:i+1]...)

		cdCmd := fmt.Sprintf("cd %s", tmpPath)
		err := sshClient.Cmd(cdCmd).Run()
		if err != nil {
			mkdirCmd := fmt.Sprintf("mkdir %s", tmpPath)
			err = sshClient.Cmd(mkdirCmd).Run()

			if err != nil {
				return err
			}
		}
	}

	return nil
}