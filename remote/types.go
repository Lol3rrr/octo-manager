package remote

import (
	ssh "github.com/helloyi/go-sshclient"
)

type session struct {
	SSHClient *ssh.Client
}
