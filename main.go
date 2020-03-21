package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"

  "octo-manager/docker"

  ssh "github.com/helloyi/go-sshclient"
)

type Config struct {
  ServerAddress string `json:"address"`
  ServerPort string `json:"port"`
  Username string `json:"username"`
  SSHKeyPath string `json:"sshkeyfile"`
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
    conf.ServerAddress + ":" + conf.ServerPort,
    conf.Username,
    conf.SSHKeyPath,
  )
  if err != nil {
    fmt.Printf("[Error] Could not connect: %v \n", err);
    return
  }
  defer client.Close()

  containers, err := docker.GetRunningContainers(client)
  if err != nil {
    fmt.Printf("[Error] Could not load Containers: %v \n", err)
    return
  }

  for _, container := range containers {
    fmt.Printf("Container: '%+v' \n", container)
  }

  return
}
