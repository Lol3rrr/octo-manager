package util

import (
	"errors"
	"io/ioutil"
	"os"
)

// GetSSHKeyPath takes a loaded Key-Path and returns the Path to the
// Key, loading it from the Environment if needed
// returns the Path, if it should be deleted and any errors that occured
func GetSSHKeyPath(loadedPath, envKeyName string) (string, bool, error) {
	if len(loadedPath) > 0 {
		return loadedPath, false, nil
	}

	SSHKey, found := os.LookupEnv(envKeyName)
	if !found || len(SSHKey) <= 0 {
		return "", false, errors.New("no SSH-Key provided in the Environment")
	}

	tmpFile, err := ioutil.TempFile(".", "ssh_*")
	if err != nil {
		return "", false, err
	}

	_, err = tmpFile.WriteString(SSHKey)
	if err != nil {
		return "", false, err
	}

	tmpFile.Close()

	return tmpFile.Name(), true, nil
}
