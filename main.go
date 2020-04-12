package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"octo-manager/auth/googledrive"
	"octo-manager/backup"
	"octo-manager/docker"
	"octo-manager/jobs"
	"octo-manager/util"

	ssh "github.com/helloyi/go-sshclient"
	"github.com/sirupsen/logrus"
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

func getEnvironment() jobs.Environment {
	result := jobs.Environment{}

	rawEnvs := os.Environ()
	for _, rawEnv := range rawEnvs {
		envContent := strings.Split(rawEnv, "=")

		key := envContent[0]
		value := envContent[1]
		result[key] = value
	}

	return result
}

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "info"
	}

	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}

	logrus.SetLevel(ll)
}

func start() int {
	configPathPtr := flag.String("config", "config.json", "The Path for the Server-Config")
	jobPathPtr := flag.String("job", "job.json", "The path for job definition")
	authNamePtr := flag.String("auth", "", "Give this a value for the part and it will run the auth stuff for it")

	flag.Parse()

	authName := *authNamePtr
	if len(authName) > 0 {
		logrus.Infof("Authing... \n")

		if authName == "googleDrive" {
			err := googledrive.Auth()
			if err != nil {
				logrus.Errorf("Could not Auth Google-Drive: %v \n", err)
			}
		} else {
			logrus.Errorf("Unknown Auth-Name: '%s' \n", authName)
		}

		return 0
	}

	conf, err := loadConfig(*configPathPtr)
	if err != nil {
		logrus.Errorf("Could not load Config: %v \n", err)
		return -1
	}

	sshKeyPath, needsDelete, err := util.GetSSHKeyPath(conf.SSHKeyPath, "SSH_KEY")
	if err != nil {
		logrus.Errorf("Loading SSH-Key-Path: %v \n", err)
		return -1
	}
	if needsDelete {
		defer os.Remove(sshKeyPath)
	}

	client, err := ssh.DialWithKey(
		conf.ServerAddress+":"+conf.ServerPort,
		conf.Username,
		sshKeyPath,
	)
	if err != nil {
		logrus.Errorf("Could not connect: %v \n", err)
		return -1
	}
	defer client.Close()

	jobSession := jobs.CreateSession(
		client,
		getEnvironment(),
		docker.NewModule(),
		backup.NewModule(),
	)

	jobConfigContent, err := ioutil.ReadFile(*jobPathPtr)
	if err != nil {
		logrus.Errorf("Could not read Job-Config: %v \n", err)
		return -1
	}

	var job jobs.Job
	err = json.Unmarshal(jobConfigContent, &job)
	if err != nil {
		logrus.Errorf("Job config file error: %v \n", err)
		return -1
	}

	err = jobSession.RunJob(&job)
	if err != nil {
		return -1
	}

	return 0
}

func main() {
	os.Exit(start())
}
