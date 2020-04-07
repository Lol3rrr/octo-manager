package jobs

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (session *Session) runStage(stage *Stage) (error, Module) {
	logrus.Infof("[Stage][%s] Starting... \n", stage.Name)

	ctx := &Ctx{
		Stage:     stage,
		SSHClient: session.SSHClient,
		Env:       session.Env,
	}

	for _, module := range session.Modules {
		if module.GetCategory() == stage.Category {
			return module.RunStage(ctx), module
		}
	}

	return fmt.Errorf("Could not find Category: '%s'", stage.Category), nil
}
