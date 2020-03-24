package backup

import ssh "github.com/helloyi/go-sshclient"

func getFilesInDir(dir string, sshClient *ssh.Client) []string {
	rawFiles, err := listFiles(dir, sshClient)
	if err != nil {
		return []string{}
	}

	result := make([]string, 0)
	for _, rawFile := range rawFiles {
		if isDir(rawFile) {
			result = append(result, getFilesInDir(rawFile, sshClient)...)
			continue
		}

		result = append(result, rawFile)
	}

	return result
}
