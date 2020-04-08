package local

import (
	"errors"
)

func getLatestDir(parentDir string) (dirInfo, error) {
	dirs, err := getBackupDirs(parentDir)
	if err != nil {
		return dirInfo{}, err
	}

	latestDir := dirInfo{}
	found := false

	for _, dir := range dirs {
		if latestDir.Timestamp < dir.Timestamp {
			latestDir = dir

			found = true
		}
	}

	if !found {
		return dirInfo{}, errors.New("could not find any Directory")
	}

	return latestDir, nil
}
