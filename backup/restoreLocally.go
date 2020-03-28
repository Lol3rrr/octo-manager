package backup

import (
	"octo-manager/backup/local"

	ssh "github.com/helloyi/go-sshclient"
)

func restoreLocally(localDir, serverDir string, sshClient *ssh.Client) error {
	local := &local.Storage{
		LocalDir: localDir,
	}

	return restore(serverDir, sshClient, local)
}
