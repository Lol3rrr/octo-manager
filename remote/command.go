package remote

import (
	"bytes"
	"io/ioutil"
	"strings"
)

func (s *session) Command(cmd string) (string, error) {
	cReader := strings.NewReader(cmd)
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

	return string(out), nil
}
