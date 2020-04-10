package general

import "strings"

// SanitizePath removes the given ServerDir string from the path
func (file *File) SanitizePath(serverDir string) {
	file.Path = strings.ReplaceAll(file.Path, serverDir+"/", "")
}
