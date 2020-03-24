package backup

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func saveFilesLocally(localDir string, files []file) error {
	if localDir[len(localDir)-1] != '/' {
		localDir += "/"
	}

	localDir += getTimestampString(time.Now().Unix()) + "/"

	for _, tmpFile := range files {
		resultPath := localDir

		tmpPath := tmpFile.Path
		if tmpPath[0] == '.' {
			tmpPath = tmpPath[1:]
		}
		if tmpPath[0] == '/' {
			tmpPath = tmpPath[1:]
		}

		resultPath += tmpPath
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
