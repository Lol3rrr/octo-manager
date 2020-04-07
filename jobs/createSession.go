package jobs

import (
	ssh "github.com/helloyi/go-sshclient"
)

// CreateSession is used to obtain an instance of a Session
func CreateSession(pSSH *ssh.Client, pEnv Environment, modules ...Module) *Session {
	return &Session{
		SSHClient: pSSH,
		Modules:   modules,
		Env:       pEnv,
	}
}
