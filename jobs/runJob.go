package jobs

import (
	"github.com/sirupsen/logrus"
)

func (session *Session) RunJob(job *Job) error {
	logrus.Infof("[Job][%s] Starting... \n", job.Name)

	for _, stage := range job.Stages {
		err, module := session.RunStage(&stage)
		if err != nil {
			logrus.Errorf("[Job][Error] %v \n", err)

			if module != nil {
				logrus.Errorf("[Job] Possible Actions: \n")
				for _, action := range module.GetActions() {
					logrus.Errorf("[Job] - '%s' \n", action)
				}
			}
			return err
		}
	}

	return nil
}
