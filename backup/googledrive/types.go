package googledrive

import (
	"octo-manager/backup/general"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

// Storage is a simple struct that holds all the needed information
type Storage struct {
	oauth2.Token
	ClientID     string
	ClientSecret string
}

type customFile struct {
	*general.File
	Dir      string
	Children []customFile
}

type backupFolder struct {
	File      *drive.File
	Timestamp int64
}
