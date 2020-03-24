package backup

import (
	ssh "github.com/helloyi/go-sshclient"
)

func getFiles(dir string, sshClient *ssh.Client) ([]file, error) {
	files := getFilesInDir(dir, sshClient)

	result := make([]file, 0, len(files))

	for _, tmpFile := range files {
		fContent := getFileContent(tmpFile, sshClient)
		result = append(result, file{
			Path:    tmpFile,
			Content: fContent,
		})
	}

	return result, nil
}
