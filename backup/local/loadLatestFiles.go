package local

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"octo-manager/backup/general"
)

func (local *Storage) LoadLatestFiles() ([]general.File, error) {
	latestDir, err := general.GetLatestDir(local.LocalDir)
	if err != nil {
		return []general.File{}, err
	}

	latestDirPath := filepath.Join(local.LocalDir, latestDir.Info.Name())

	files := make([]general.File, 0)
	err = filepath.Walk(latestDirPath, func(rawPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(rawPath)
		if err != nil {
			return nil
		}

		path := strings.ReplaceAll(rawPath, latestDirPath+"/", "")

		files = append(files, general.File{
			Path:    path,
			Content: string(content),
		})

		return nil
	})

	return files, nil
}
