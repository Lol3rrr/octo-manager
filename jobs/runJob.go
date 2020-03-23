package jobs

import (
	"fmt"
)

func (session *Session) RunJob(job *Job) error {
	fmt.Printf("[Job][%s] Starting... \n", job.Name)

	for _, stage := range job.Stages {
		err := session.RunStage(&stage)
		if err != nil {
			fmt.Printf("[Job][Error] Failed '%v' \n", err)
			return err
		}
	}

	return nil
}
