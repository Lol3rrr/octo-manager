package remote

import "fmt"

func (s *session) GetFileContent(path string) string {
	commandString := fmt.Sprintf("cat %s", path)
	rawOutput, err := s.Command(commandString)
	if err != nil {
		return ""
	}

	return string(rawOutput)
}
