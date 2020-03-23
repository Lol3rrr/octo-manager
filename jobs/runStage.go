package jobs

import (
	"fmt"
)

func (session *Session) RunStage(stage *Stage) error {
	fmt.Printf("[Stage][%s] Starting... \n", stage.Name)

	for _, module := range session.Modules {
		if module.GetCategory() == stage.Category {
			return module.RunStage(stage, session.SSHClient, session.Env)
		}
	}

	return fmt.Errorf("Could not find Category: '%s'", stage.Category)
}
