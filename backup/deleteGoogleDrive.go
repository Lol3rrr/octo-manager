package backup

import (
	"errors"
	"octo-manager/backup/googledrive"
	"octo-manager/jobs"
	"octo-manager/remote"

	"golang.org/x/oauth2"
)

func deleteGoogleDrive(stage *jobs.Stage, env jobs.Environment, remoteCon remote.Session) error {
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

	return delete(stage, env, drive)
}