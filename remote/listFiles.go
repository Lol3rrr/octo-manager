package remote

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func (s *session) listFiles(dir string) ([]string, error) {
	if len(dir) == 0 {
		dir = "."
	}
	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}

	logrus.Debugf("[ListFiles] Listing Dir: '%s' \n", dir)

	commandString := fmt.Sprintf("ls -a -p %s/", dir)
	rawOutput, err := s.Command(commandString)
	if err != nil {
		return []string{}, err
	}

	output := string(rawOutput)
	outputLines := strings.Split(output, "\n")
	outputLines = outputLines[:len(outputLines)-1]

	result := make([]string, 0)
	for _, line := range outputLines {
		if len(line) <= 0 || line == "./" || line == "../" {
			continue
		}

		path := dir + "/" + line

		logrus.Debugf("[ListFiles] Found Path: '%s' \n", path)
		result = append(result, path)
	}

	return result, nil
}
