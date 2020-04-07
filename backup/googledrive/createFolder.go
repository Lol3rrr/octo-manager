package googledrive

import (
	"google.golang.org/api/drive/v3"
)

func createFolder(service *drive.Service, name, parentID string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: dirMimeType,
		Parents:  []string{parentID},
	}

	return service.Files.Create(d).EnforceSingleParent(true).Do()
}
