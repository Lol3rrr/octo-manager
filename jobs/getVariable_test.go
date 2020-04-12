package jobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariable(t *testing.T) {
	tables := []struct {
		Name        string
		InputStage  Stage
		InputKey    string
		InputEnv    Environment
		OutputValue string
		OutputFound bool
	}{
		{
			Name: "Valid Input, key in Variables",
			InputStage: Stage{
				Variables: map[string]string{
					"testKey": "testValue",
				},
			},
			InputKey:    "testKey",
			InputEnv:    Environment{},
			OutputValue: "testValue",
			OutputFound: true,
		},
		{
			Name:       "Valid Input, key in Variables",
			InputStage: Stage{},
			InputKey:   "testKey",
			InputEnv: Environment{
				"testKey": "testValue",
			},
			OutputValue: "testValue",
			OutputFound: true,
		},
		{
			Name:        "Invalid Input, key not found",
			InputStage:  Stage{},
			InputKey:    "testKey",
			InputEnv:    Environment{},
			OutputValue: "",
			OutputFound: false,
		},
	}

	for _, table := range tables {
		inStage := table.InputStage
		inKey := table.InputKey
		inEnv := table.InputEnv
		outValue := table.OutputValue
		outFound := table.OutputFound

		t.Run(table.Name, func(t *testing.T) {
			resultValue, resultFound := inStage.GetVariable(inKey, inEnv)

			assert.Equal(t, outValue, resultValue)
			assert.Equal(t, outFound, resultFound)
		})
	}
}
