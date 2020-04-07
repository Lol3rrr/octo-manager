package remote

import (
	ssh "github.com/helloyi/go-sshclient"
)

// CreateRemoteCon initializes a remote-Con session using the SSH-Session
func CreateRemoteCon(sshClient *ssh.Client) Session {
	return &session{
		SSHClient: sshClient,
	}
}
