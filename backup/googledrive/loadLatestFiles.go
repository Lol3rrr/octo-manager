package googledrive

import (
	"errors"
	"fmt"
	"io/ioutil"
	"octo-manager/auth/googledrive"
	"octo-manager/backup/general"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
)

// LoadLatestFiles is used to get the Files from the latest Backup folder
func (storage *Storage) LoadLatestFiles() ([]general.File, error) {
	service, err := googledrive.GetDrive(storage.ClientID, storage.ClientSecret, &storage.Token)
	if err != nil {
		return []general.File{}, err
	}

	logrus.Infof("Getting a list of all Files... \n")

	fileList, err := service.Files.List().Do()
	if err != nil {
		return []general.File{}, err
	}

	rawFiles := fileList.Files

	for len(fileList.NextPageToken) > 0 {
		fileList, err = service.Files.List().PageToken(fileList.NextPageToken).Do()
		if err != nil {
			break
		}

		rawFiles = append(rawFiles, fileList.Files...)
	}

	logrus.Infof("Getting the latest Folder... \n")

	var latestFile *drive.File = nil
	latestTimestamp := int64(-1)
	for _, file := range rawFiles {
		if file.MimeType != dirMimeType {
			continue
		}

		timestamp := general.GetTimestampFromString(file.Name)
		if timestamp <= latestTimestamp {
			continue
		}

		latestFile = file
		latestTimestamp = timestamp
	}

	if latestFile == nil {
		return []general.File{}, errors.New("could not find a fitting file")
	}

	logrus.Infof("Loading the latest Files... \n")

	query := fmt.Sprintf("parents in '%s'", latestFile.Id)
	filesInFolder, err := service.Files.List().Q(query).Do()
	if err != nil {
		return []general.File{}, err
	}

	rawResult := filesInFolder.Files

	for len(filesInFolder.NextPageToken) > 0 {
		filesInFolder, err = service.Files.List().Q(query).PageToken(filesInFolder.NextPageToken).Do()
		if err != nil {
			break
		}

		rawResult = append(rawResult, filesInFolder.Files...)
	}

	logrus.Infof("Downloading files... \n")

	result := make([]general.File, 0, len(rawResult))
	for _, rawFile := range rawResult {
		response, err := service.Files.Get(rawFile.Id).Download()
		if err != nil {
			logrus.Errorf("Could not download File: %v \n", err)
			continue
		}

		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("Could not read File Content: %v \n", err)
			response.Body.Close()
			continue
		}

		response.Body.Close()

		result = append(result, general.File{
			Path:    rawFile.Name,
			Content: string(content),
		})
	}

	return result, nil
}
