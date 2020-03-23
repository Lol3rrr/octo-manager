package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullImage(t *testing.T) {
	tables := []struct {
		Name        string
		InputImage  string
		ResultError bool
	}{
		{
			Name:        "Invalid Input, empty Image",
			InputImage:  "",
			ResultError: true,
		},
	}

	for _, table := range tables {
		inImage := table.InputImage
		resultError := table.ResultError

		t.Run(table.Name, func(t *testing.T) {
			outputError := pullImage(nil, inImage)

			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
