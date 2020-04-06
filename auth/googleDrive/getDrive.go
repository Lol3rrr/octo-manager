package googleDrive

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

// GetDrive returns a new Drive using the given Data
func GetDrive(clientID, clientSecret string, tok *oauth2.Token) (*drive.Service, error) {
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: 0,
		},
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		Scopes: []string{
			drive.DriveFileScope,
		},
	}

	client := config.Client(context.Background(), tok)

	return drive.New(client)
}
