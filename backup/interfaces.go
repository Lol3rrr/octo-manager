package backup

import "octo-manager/backup/general"

type storage interface {
	Save(string, []general.File) error
	LoadLatestFiles() ([]general.File, error)
}
