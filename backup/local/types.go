package local

import "os"

// Storage is a simple struct for local Structure
type Storage struct {
	LocalDir string
}

type dirInfo struct {
	Timestamp int64
	Info      os.FileInfo
}
