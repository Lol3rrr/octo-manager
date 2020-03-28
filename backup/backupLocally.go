package backup

import (
	"octo-manager/backup/local"

	ssh "github.com/helloyi/go-sshclient"
)

func backupLocally(serverDir, localDir string, sshClient *ssh.Client) error {
	local := &local.Storage{
		LocalDir: localDir,
	}

	return backup(serverDir, sshClient, local)
}
