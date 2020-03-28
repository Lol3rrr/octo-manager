package remote

import (
	"octo-manager/backup/general"

	ssh "github.com/helloyi/go-sshclient"
)

func GetFiles(dir string, sshClient *ssh.Client) ([]general.File, error) {
	files := getFilesInDir(dir, sshClient)

	result := make([]general.File, 0, len(files))

	for _, tmpFile := range files {
		fContent := getFileContent(tmpFile, sshClient)
		result = append(result, general.File{
			Path:    tmpFile,
			Content: fContent,
		})
	}

	return result, nil
}
