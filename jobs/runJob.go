package jobs

import (
	"fmt"
)

func (session *Session) RunJob(job *Job, env Environment) error {
	fmt.Printf("[Job][%s] Starting... \n", job.Name)

	for _, stage := range job.Stages {
		err := stage.Run(session.SSHClient, env)
		if err != nil {
			fmt.Printf("[Job][Error] Failed '%v' \n", err)
			return err
		}
	}

	return nil
}
