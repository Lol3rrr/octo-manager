package remote

import (
	"octo-manager/remote"
)

func getFilesInDir(dir string, remoteCon remote.Session) []string {
	rawFiles, err := listFiles(dir, remoteCon)
	if err != nil {
		return []string{}
	}

	result := make([]string, 0)
	for _, rawFile := range rawFiles {
		if isDir(rawFile) {
			result = append(result, getFilesInDir(rawFile, remoteCon)...)
			continue
		}

		result = append(result, rawFile)
	}

	return result
}
