package local

import (
	"fmt"
	"io/ioutil"
	"octo-manager/backup/general"
	"os"
	"path/filepath"
	"time"
)

// Save is used to save the given files locally
func (local *Storage) Save(files []general.File) error {
	if local.LocalDir[len(local.LocalDir)-1] != '/' {
		local.LocalDir += "/"
	}

	local.LocalDir += general.GetTimestampString(time.Now().Unix()) + "/"

	for _, tmpFile := range files {
		resultPath := local.LocalDir

		resultPath += tmpFile.Path
		resultDir := filepath.Dir(resultPath)

		err := os.MkdirAll(resultDir, os.ModePerm)
		if err != nil {
			fmt.Printf("[Error] %v \n", err)
		}

		err = ioutil.WriteFile(resultPath, []byte(tmpFile.Content), 0644)
		if err != nil {
			fmt.Printf("[Error] %v \n", err)
		}
	}

	return nil
}
