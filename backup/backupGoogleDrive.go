package backup

import (
	"errors"
	"octo-manager/backup/googleDrive"
	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
	"golang.org/x/oauth2"
)

func backupGoogleDrive(stage *jobs.Stage, env jobs.Environment, sshClient *ssh.Client) error {
	serverDir, found := stage.GetVariable("serverDir", env)
	if !found {
		return errors.New("Missing Variable: 'serverDir'")
	}
	refreshToken, found := stage.GetVariable("refreshToken", env)
	if !found {
		return errors.New("Missing Variable: 'refreshToken'")
	}
	clientID, found := stage.GetVariable("clientID", env)
	if !found {
		return errors.New("Missing Variable: 'clientID'")
	}
	clientSecret, found := stage.GetVariable("clientSecret", env)
	if !found {
		return errors.New("Missing Variable: 'clientSecret'")
	}

	drive := &googleDrive.Storage{
		Token: oauth2.Token{
			TokenType:    "Bearer",
			RefreshToken: refreshToken,
		},
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return backup(serverDir, sshClient, drive)
}
