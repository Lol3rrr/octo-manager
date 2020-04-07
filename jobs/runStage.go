package jobs

import (
	"fmt"
	"octo-manager/remote"

	"github.com/sirupsen/logrus"
)

func (session *Session) runStage(stage *Stage) (error, Module) {
	logrus.Infof("[Stage][%s] Starting... \n", stage.Name)

	ctx := &Ctx{
		Stage:     stage,
		Env:       session.Env,
		RemoteCon: remote.CreateRemoteCon(session.SSHClient),
	}

	for _, module := range session.Modules {
		if module.GetCategory() == stage.Category {
			return module.RunStage(ctx), module
		}
	}

	return fmt.Errorf("Could not find Category: '%s'", stage.Category), nil
}
