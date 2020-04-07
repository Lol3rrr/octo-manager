package googledrive

import (
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
)

// Auth is used to authorize and generate new Token for Google-Drive
func Auth() error {
	credPath := "googleDriveCredentials.json"
	b, err := ioutil.ReadFile(credPath)
	if err != nil {
		return err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return err
	}

	tok, err := getTokenFromWeb(config)
	if err != nil {
		return err
	}

	saveToken("googleDriveToken.json", tok)

	return nil
}
