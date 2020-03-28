package general

import "time"

func GetTimestampString(timestamp int64) string {
	tmpTime := time.Unix(timestamp, 0)

	return tmpTime.Format("2006-01-02_15-04")
}

func getTimestampFromString(timestamp string) int64 {
	rawTime, err := time.Parse("2006-01-02_15-04", timestamp)
	if err != nil {
		return -1
	}

	return rawTime.Unix()
}
