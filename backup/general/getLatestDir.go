package general

import (
	"errors"
	"io/ioutil"
	"os"
)

type dirInfo struct {
	Timestamp int64
	Info      os.FileInfo
}

func GetLatestDir(parentDir string) (dirInfo, error) {
	fileInfos, err := ioutil.ReadDir(parentDir)
	if err != nil {
		return dirInfo{}, err
	}

	latestDir := dirInfo{
		Timestamp: -1,
		Info:      nil,
	}
	found := false

	for _, fInfo := range fileInfos {
		if fInfo.IsDir() {
			timestamp := GetTimestampFromString(fInfo.Name())
			if timestamp < 0 {
				continue
			}

			if latestDir.Timestamp < timestamp {
				latestDir = dirInfo{
					Timestamp: timestamp,
					Info:      fInfo,
				}

				found = true
			}
		}
	}

	if !found {
		return dirInfo{}, errors.New("Could not find any Directory")
	}

	return latestDir, nil
}
