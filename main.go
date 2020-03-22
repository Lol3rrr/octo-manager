package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"octo-manager/jobs"

	ssh "github.com/helloyi/go-sshclient"
)

type Config struct {
	ServerAddress string `json:"address"`
	ServerPort    string `json:"port"`
	Username      string `json:"username"`
	SSHKeyPath    string `json:"sshkeyfile"`
}

func loadConfig() (*Config, error) {
	configContent, err := ioutil.ReadFile("./config.json")
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
	fmt.Printf("Starting... \n")

	conf, err := loadConfig()
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

	jobPathPtr := flag.String("job", "job.json", "The path for job definition")

	flag.Parse()

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
	fmt.Printf("Jobs Error: '%v' \n", err)

	return
}
