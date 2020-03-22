package jobs

import (
	ssh "github.com/helloyi/go-sshclient"
)

func CreateSession(pSSH *ssh.Client, modules ...Module) *Session {
	return &Session{
		SSHClient: pSSH,
		Modules:   modules,
	}
}
