package remote

import (
	"fmt"
	"octo-manager/remote"
)

func getFileContent(path string, remoteCon remote.Session) string {
	commandString := fmt.Sprintf("cat %s", path)
	rawOutput, err := remoteCon.Command(commandString)
	if err != nil {
		return ""
	}

	return string(rawOutput)
}
