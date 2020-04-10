package remote

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func (s *session) ShellCommand(cmd string) (string, error) {
	rawSeperator := make([]byte, 100)
	rand.Read(rawSeperator)
	seperator := base64.StdEncoding.EncodeToString(rawSeperator)
	preCmd := fmt.Sprintf("echo '%s';", seperator)

	cReader := strings.NewReader(preCmd + cmd)
	cWriter := bytes.NewBuffer([]byte{})
	shell := s.SSHClient.Shell().SetStdio(cReader, cWriter, cWriter)
	err := shell.Start()
	if err != nil {
		return "", err
	}

	out, err := ioutil.ReadAll(cWriter)
	if err != nil {
		return "", err
	}

	rawOutStr := string(out)
	seperatorIndex := strings.Index(rawOutStr, seperator)
	if seperatorIndex < 0 {
		return "", errors.New("command was not executed correctly")
	}

	sectionStart := seperatorIndex + len(seperator)

	return rawOutStr[sectionStart:], nil
}
