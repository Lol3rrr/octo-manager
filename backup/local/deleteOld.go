package local

import (
	"os"
	"path/filepath"
)

// DeleteOld is used to delete all backup folders older than Threshold
func (s *Storage) DeleteOld(thresholdTime int64) error {
	dirs, err := getBackupDirs(s.LocalDir)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if thresholdTime < dir.Timestamp {
			continue
		}

		path := filepath.Join(s.LocalDir, dir.Info.Name())
		os.RemoveAll(path)
	}

	return nil
}
