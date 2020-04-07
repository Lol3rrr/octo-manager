package remote

import "octo-manager/backup/general"

// Session is a simple interface that represents a
// connection to a server over ssh
type Session interface {
	// Command simply executes a command on the Server and
	// returns the result
	Command(cmd string) (string, error)
	// GetFileContent returns the Content of the file for the given path
	// Empty if an error occured
	GetFileContent(path string) string
	// WriteFile simply writes the given content to the provided path
	// returns an error, if one occured
	WriteFile(path, content string) error
	// GetFilesInDir simply returns all the paths to files in the given
	// directory and its sub-directorys
	GetFilesInDir(dir string) []string
	// WriteFiles simply writes all the files into the given
	// directory on the Server
	// returns an error if any occured
	WriteFiles(files []general.File, serverDir string) error
	// GetFiles loads all files and their content from the given
	// directory
	// returns the files and any error that may have occured
	GetFiles(dir string) ([]general.File, error)
}
