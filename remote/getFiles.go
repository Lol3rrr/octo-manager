package remote

import (
	"octo-manager/backup/general"

	"github.com/sirupsen/logrus"
)

func (s *session) GetFiles(dir string) ([]general.File, error) {
	logrus.Infof("[GetFiles] Loading File paths... \n")
	files := s.GetFilesInDir(dir)

	result := make([]general.File, 0, len(files))

	logrus.Infof("[GetFiles] Loading the Files content... \n")
	for _, tmpFile := range files {
		fContent := s.GetFileContent(tmpFile)
		result = append(result, general.File{
			Path:    tmpFile,
			Content: fContent,
		})
	}

	return result, nil
}
