package remote

import (
	"bytes"
	"fmt"
	"strings"

	ssh "github.com/helloyi/go-sshclient"
)

func writeToFile(content, path string, sshClient *ssh.Client) error {
	writeCmd := fmt.Sprintf("cat > %s \n", path)

	readContent := writeCmd + content
	cReader := strings.NewReader(readContent)
	cWriter := bytes.NewBuffer([]byte{})
	shell := sshClient.Shell().SetStdio(cReader, cWriter, cWriter)
	err := shell.Start()
	if err != nil {
		return err
	}

	return nil
}
