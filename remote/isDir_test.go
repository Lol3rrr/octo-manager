package remote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	tables := []struct {
		Name      string
		InputPath string
		Output    bool
	}{
		{
			Name:      "Valid Input, dir",
			InputPath: "test/",
			Output:    true,
		},
		{
			Name:      "Valid Input, file",
			InputPath: "test",
			Output:    false,
		},
		{
			Name:      "Empty path",
			InputPath: "",
			Output:    false,
		},
	}

	for _, table := range tables {
		inPath := table.InputPath
		output := table.Output

		t.Run(table.Name, func(t *testing.T) {
			result := isDir(inPath)

			assert.Equal(t, output, result)
		})
	}
}
