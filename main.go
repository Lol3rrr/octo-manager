package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
)

// Config is a simple struct that holds the basic server config
type Config struct {
	ServerAddress string `json:"address"`
	ServerPort    string `json:"port"`
	Username      string `json:"username"`
	SSHKeyPath    string `json:"sshkeyfile"`
}

func loadConfig(path string) (*Config, error) {
	configContent, err := ioutil.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var result Config
	err = json.Unmarshal(configContent, &result)
	if err != nil {
		return &Config{}, err
	}

	return &result, nil
}

func main() {
	configPathPtr := flag.String("config", "config.json", "The Path for the Server-Config")
	jobPathPtr := flag.String("job", "job.json", "The path for job definition")

	flag.Parse()

	conf, err := loadConfig(*configPathPtr)
	if err != nil {
		fmt.Printf("[Error] Could not load Config: %v \n", err)
		return
	}

	client, err := ssh.DialWithKey(
		conf.ServerAddress+":"+conf.ServerPort,
		conf.Username,
		conf.SSHKeyPath,
	)
	if err != nil {
		fmt.Printf("[Error] Could not connect: %v \n", err)
		return
	}
	defer client.Close()

	jobSession := jobs.CreateSession(client)

	jobConfigContent, err := ioutil.ReadFile(*jobPathPtr)
	if err != nil {
		fmt.Printf("[Error] Could not read Job-Config: %v \n", err)
		return
	}

	var job jobs.Job
	err = json.Unmarshal(jobConfigContent, &job)

	err = jobSession.RunJob(&job, jobs.Environment{
		"image": "lol3r/personal_site:latest",
	})

	if err != nil {
		fmt.Printf("Jobs Error: '%v' \n", err)
	}

	return
}
