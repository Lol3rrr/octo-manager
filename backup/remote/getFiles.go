package remote

import (
	"octo-manager/backup/general"
	"octo-manager/remote"
)

func GetFiles(dir string, remoteCon remote.Session) ([]general.File, error) {
	files := getFilesInDir(dir, remoteCon)

	result := make([]general.File, 0, len(files))

	for _, tmpFile := range files {
		fContent := getFileContent(tmpFile, remoteCon)
		result = append(result, general.File{
			Path:    tmpFile,
			Content: fContent,
		})
	}

	return result, nil
}
