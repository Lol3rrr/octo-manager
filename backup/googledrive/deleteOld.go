package googledrive

import (
	"octo-manager/auth/googledrive"
)

// DeleteOld deletes all backup folder older than the treshold
func (storage *Storage) DeleteOld(thresholdTime int64) error {
	service, err := googledrive.GetDrive(storage.ClientID, storage.ClientSecret, &storage.Token)
	if err != nil {
		return err
	}

	backupDirs, err := getBackupDirs(service)
	if err != nil {
		return err
	}

	for _, dir := range backupDirs {
		if thresholdTime > dir.Timestamp {
			continue
		}

		dir.Delete(service)
	}

	return nil
}
