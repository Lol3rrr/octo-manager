package googledrive

import (
	"octo-manager/backup/general"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
)

func getBackupDirs(service *drive.Service) ([]backupFolder, error) {
	logrus.Infof("Getting a list of all Files... \n")

	fileList, err := service.Files.List().Do()
	if err != nil {
		return []backupFolder{}, err
	}

	rawFiles := fileList.Files

	for len(fileList.NextPageToken) > 0 {
		fileList, err = service.Files.List().PageToken(fileList.NextPageToken).Do()
		if err != nil {
			break
		}

		rawFiles = append(rawFiles, fileList.Files...)
	}

	result := make([]backupFolder, 0)

	for _, rawFile := range rawFiles {
		if rawFile.MimeType != dirMimeType {
			continue
		}

		timestamp := general.GetTimestampFromString(rawFile.Name)
		if timestamp < 0 {
			continue
		}

		result = append(result, backupFolder{
			File:      rawFile,
			Timestamp: timestamp,
		})
	}

	return result, nil
}
