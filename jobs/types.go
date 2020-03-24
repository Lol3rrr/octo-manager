package jobs

import (
	ssh "github.com/helloyi/go-sshclient"
)

// Session represents a single Session with all its settings
type Session struct {
	Modules   []Module
	SSHClient *ssh.Client
	Env       Environment
}

// Module is a simple abstraction, that represents a single Category
type Module interface {
	GetCategory() string
	RunStage(*Stage, *ssh.Client, Environment) error
	GetActions() []string
}

// Job defines a single Job, like deployment, and includes
// all Details needed for this job
type Job struct {
	Name   string  `json:"name"`
	Stages []Stage `json:"stages"`
}

// Stage defines a single Stage of a Job
type Stage struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    string            `json:"category"`
	Action      string            `json:"action"`
	Variables   map[string]string `json:"variables"`
}

// Environment simply contains all the given Env-Variables
type Environment map[string]string
