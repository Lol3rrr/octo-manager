package backup

import (
	"errors"
	"octo-manager/jobs"
	"strconv"
	"time"
)

func delete(stage *jobs.Stage, env jobs.Environment, storageInterface storage) error {
	rawThreshold, found := stage.GetVariable("threshold", env)
	if !found {
		return errors.New("missing Variable 'threshold'")
	}

	threshold, err := strconv.Atoi(rawThreshold)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	treshHoldTime := now - int64(threshold*60*60)

	return storageInterface.DeleteOld(treshHoldTime)
}
