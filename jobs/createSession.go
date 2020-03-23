package jobs

import (
	ssh "github.com/helloyi/go-sshclient"
)

func CreateSession(pSSH *ssh.Client, pEnv Environment, modules ...Module) *Session {
	return &Session{
		SSHClient: pSSH,
		Modules:   modules,
		Env:       pEnv,
	}
}
