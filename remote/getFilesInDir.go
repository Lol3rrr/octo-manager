package remote

func (s *session) GetFilesInDir(dir string) []string {
	rawFiles, err := s.listFiles(dir)
	if err != nil {
		return []string{}
	}

	result := make([]string, 0)
	for _, rawFile := range rawFiles {
		if isDir(rawFile) {
			result = append(result, s.GetFilesInDir(rawFile)...)
			continue
		}

		result = append(result, rawFile)
	}

	return result
}