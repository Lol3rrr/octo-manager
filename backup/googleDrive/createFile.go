package googleDrive

import (
	"io"

	"google.golang.org/api/drive/v3"
)

func createFile(service *drive.Service, parentID, name string, content io.Reader) (*drive.File, error) {
	f := &drive.File{
		Name:     name,
		MimeType: "text",
		Parents:  []string{parentID},
	}

	return service.Files.Create(f).Media(content).Do()
}
