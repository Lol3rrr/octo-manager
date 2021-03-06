package backup

import "octo-manager/backup/general"

type storage interface {
	Save([]general.File) error
	LoadLatestFiles() ([]general.File, error)
	DeleteOld(thresholdTime int64) error
}
