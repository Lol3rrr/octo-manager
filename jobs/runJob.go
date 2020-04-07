package jobs

import (
	"github.com/sirupsen/logrus"
)

// RunJob is used to run any given Job from the Config
func (session *Session) RunJob(job *Job) error {
	logrus.Infof("[Job][%s] Starting... \n", job.Name)

	for _, stage := range job.Stages {
		module, err := session.runStage(&stage)
		if err != nil {
			logrus.Errorf("[Job] %v \n", err)

			if module != nil {
				actions := module.GetActions()
				if len(actions) > 0 {
					logrus.Errorf("[Job] Possible Actions: \n")
					for _, action := range module.GetActions() {
						logrus.Errorf("[Job] - '%s' \n", action)
					}
				}
			}
			return err
		}
	}

	return nil
}
