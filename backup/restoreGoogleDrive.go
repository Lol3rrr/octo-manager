package backup

import (
	"errors"
	"octo-manager/backup/googledrive"
	"octo-manager/jobs"
	"octo-manager/remote"

	"golang.org/x/oauth2"
)

func restoreGoogleDrive(stage *jobs.Stage, env jobs.Environment, remoteCon remote.Session) error {
	serverDir, found := stage.GetVariable("serverDir", env)
	if !found {
		return errors.New("missing Variable: 'serverDir'")
	}
	refreshToken, found := stage.GetVariable("refreshToken", env)
	if !found {
		return errors.New("missing Variable: 'refreshToken'")
	}
	clientID, found := stage.GetVariable("clientID", env)
	if !found {
		return errors.New("missing Variable: 'clientID'")
	}
	clientSecret, found := stage.GetVariable("clientSecret", env)
	if !found {
		return errors.New("missing Variable: 'clientSecret'")
	}

	drive := &googledrive.Storage{
		Token: oauth2.Token{
			TokenType:    "Bearer",
			RefreshToken: refreshToken,
		},
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return restore(serverDir, remoteCon, drive)
}
