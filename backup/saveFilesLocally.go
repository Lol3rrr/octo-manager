package backup

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func saveFilesLocally(serverDir, localDir string, files []file) error {
	if localDir[len(localDir)-1] != '/' {
		localDir += "/"
	}

	localDir += getTimestampString(time.Now().Unix()) + "/"

	for _, tmpFile := range files {
		resultPath := localDir

		resultPath += strings.ReplaceAll(tmpFile.Path, serverDir, "")
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
