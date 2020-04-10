package remote

import "fmt"

func (s *session) GetFileContent(path string) string {
	commandString := fmt.Sprintf("cat %s", path)
	output, err := s.Command(commandString)
	if err != nil {
		return ""
	}

	return output
}
