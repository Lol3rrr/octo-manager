package docker

import (
	"strings"

	ssh "github.com/helloyi/go-sshclient"
)

func getRunningContainers(sshClient *ssh.Client) ([]Container, error) {
	outputRaw, err := sshClient.Cmd("docker ps").Output()
	if err != nil {
		return []Container{}, err
	}

	lines := strings.Split(string(outputRaw), "\n")
	lines = lines[1 : len(lines)-1]
	result := make([]Container, 0, len(lines))

	for _, line := range lines {
		container, err := parseContainerLine(line)
		if err != nil {
			continue
		}

		result = append(result, container)
	}

	return result, nil
}
