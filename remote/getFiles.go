package remote

import "octo-manager/backup/general"

func (s *session) GetFiles(dir string) ([]general.File, error) {
	files := s.GetFilesInDir(dir)

	result := make([]general.File, 0, len(files))

	for _, tmpFile := range files {
		fContent := s.GetFileContent(tmpFile)
		result = append(result, general.File{
			Path:    tmpFile,
			Content: fContent,
		})
	}

	return result, nil
}
