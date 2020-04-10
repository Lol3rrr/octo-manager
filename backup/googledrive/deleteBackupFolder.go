package googledrive

import "google.golang.org/api/drive/v3"

func (folder *backupFolder) Delete(service *drive.Service) error {
	return service.Files.Delete(folder.File.Id).Do()
}
