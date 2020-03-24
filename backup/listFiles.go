package backup

import (
	"fmt"
	"strings"

	ssh "github.com/helloyi/go-sshclient"
)

func listFiles(dir string, sshClient *ssh.Client) ([]string, error) {
	if len(dir) == 0 {
		dir = "."
	}
	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}

	commandString := fmt.Sprintf("ls -a -p %s/", dir)
	rawOutput, err := sshClient.Cmd(commandString).Output()
	if err != nil {
		return []string{}, err
	}

	output := string(rawOutput)
	outputLines := strings.Split(output, "\n")
	outputLines = outputLines[:len(outputLines)-1]

	result := make([]string, 0, len(outputLines)-2)
	for _, line := range outputLines {
		if line == "./" || line == "../" {
			continue
		}
		result = append(result, dir+"/"+line)
	}

	return result, nil
}
