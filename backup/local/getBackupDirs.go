package local

import (
	"io/ioutil"
	"octo-manager/backup/general"
)

// getBackupDirs returns all the directorys that correspond to
// backups in the given path
func getBackupDirs(path string) ([]dirInfo, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return []dirInfo{}, err
	}

	dirs := make([]dirInfo, 0)
	for _, fInfo := range fileInfos {
		if !fInfo.IsDir() {
			continue
		}

		timestamp := general.GetTimestampFromString(fInfo.Name())
		if timestamp < 0 {
			continue
		}

		dirs = append(dirs, dirInfo{
			Timestamp: timestamp,
			Info:      fInfo,
		})
	}

	return dirs, nil
}
