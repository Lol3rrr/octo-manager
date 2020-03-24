package backup

import "time"

func getTimestampString(timestamp int64) string {
	tmpTime := time.Unix(timestamp, 0)

	return tmpTime.Format("2006-01-02_15-04")
}
