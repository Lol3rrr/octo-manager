package googledrive

import (
	"octo-manager/auth/googledrive"
	"octo-manager/backup/general"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Save is used to save the given files to google Drive
func (storage *Storage) Save(files []general.File) error {
	folderName := general.GetTimestampString(time.Now().Unix())

	service, err := googledrive.GetDrive(storage.ClientID, storage.ClientSecret, &storage.Token)
	if err != nil {
		return err
	}

	backupRoot, err := createFolder(service, folderName, "root")
	if err != nil {
		return err
	}

	for _, file := range files {
		_, err = createFile(service, backupRoot.Id, file.Path, strings.NewReader(file.Content))
		if err != nil {
			logrus.Errorf("Could not create File: %v \n", err)
			continue
		}
	}

	return nil
}
