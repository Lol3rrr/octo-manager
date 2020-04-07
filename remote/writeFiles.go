package remote

import (
	"octo-manager/backup/general"
	"path/filepath"
)

func (s *session) WriteFiles(files []general.File, serverDir string) error {
	for _, tmpFile := range files {
		tmpPath := filepath.Join(serverDir, tmpFile.Path)
		err := s.mkdirs(filepath.Dir(tmpPath))
		if err != nil {
			return err
		}

		err = s.WriteFile(tmpPath, tmpFile.Content)
		if err != nil {
			return err
		}
	}

	return nil
}
