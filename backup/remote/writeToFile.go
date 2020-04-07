package remote

import (
	"fmt"
	"octo-manager/remote"
)

func writeToFile(content, path string, remoteCon remote.Session) error {
	writeCmd := fmt.Sprintf("cat > %s \n", path)

	readContent := writeCmd + content
	_, err := remoteCon.Command(readContent)
	if err != nil {
		return err
	}

	return nil
}
