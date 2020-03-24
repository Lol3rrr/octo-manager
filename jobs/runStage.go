package jobs

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (session *Session) RunStage(stage *Stage) (error, Module) {
	logrus.Infof("[Stage][%s] Starting... \n", stage.Name)

	for _, module := range session.Modules {
		if module.GetCategory() == stage.Category {
			return module.RunStage(stage, session.SSHClient, session.Env), module
		}
	}

	return fmt.Errorf("Could not find Category: '%s'", stage.Category), nil
}
