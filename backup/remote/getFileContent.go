package remote

import (
	"fmt"

	ssh "github.com/helloyi/go-sshclient"
)

func getFileContent(path string, sshClient *ssh.Client) string {
	commandString := fmt.Sprintf("cat %s", path)
	rawOutput, err := sshClient.Cmd(commandString).Output()
	if err != nil {
		return ""
	}

	return string(rawOutput)
}
