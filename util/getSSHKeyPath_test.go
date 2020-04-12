package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSSHKeyPath(t *testing.T) {
	tables := []struct {
		Name             string
		EnvKey           string
		EnvValue         string
		InputLoadedPath  string
		InputEnvKeyName  string
		OutputPath       string
		OutputRandomPath bool
		OutputDelete     bool
		OutputError      bool
	}{
		{
			Name:             "Valid Input, should just return the loadedPath",
			EnvKey:           "SSH_KEY",
			EnvValue:         "TestSSHKey",
			InputLoadedPath:  "./test/key",
			InputEnvKeyName:  "SSH_KEY",
			OutputPath:       "./test/key",
			OutputRandomPath: false,
			OutputDelete:     false,
			OutputError:      false,
		},
		{
			Name:             "Valid Input, should load from the Env-Value",
			EnvKey:           "SSH_KEY",
			EnvValue:         "TestSSHKey",
			InputLoadedPath:  "",
			InputEnvKeyName:  "SSH_KEY",
			OutputPath:       "",
			OutputRandomPath: true,
			OutputDelete:     true,
			OutputError:      false,
		},
		{
			Name:             "Invalid Input, Env-Value is empty",
			EnvKey:           "SSH_KEY",
			EnvValue:         "",
			InputLoadedPath:  "",
			InputEnvKeyName:  "SSH_KEY",
			OutputPath:       "",
			OutputRandomPath: true,
			OutputDelete:     false,
			OutputError:      true,
		},
		{
			Name:             "Invalid Input, Env-Value is not set",
			EnvKey:           "SSH_KEY",
			EnvValue:         "./test/key",
			InputLoadedPath:  "",
			InputEnvKeyName:  "WRONG_SSH_KEY",
			OutputPath:       "",
			OutputRandomPath: true,
			OutputDelete:     false,
			OutputError:      true,
		},
	}

	for _, table := range tables {
		envKey := table.EnvKey
		envValue := table.EnvValue
		inLoadedPath := table.InputLoadedPath
		inEnvKeyName := table.InputEnvKeyName
		outPath := table.OutputPath
		outRandomPath := table.OutputRandomPath
		outDelete := table.OutputDelete
		outError := table.OutputError

		t.Run(table.Name, func(t *testing.T) {
			os.Setenv(envKey, envValue)

			resultPath, resultDelete, resultErr := GetSSHKeyPath(inLoadedPath, inEnvKeyName)
			if resultDelete && len(resultPath) > 0 {
				defer os.Remove(resultPath)
			}

			if outError {
				assert.NotNil(t, resultErr)
			} else {
				assert.Nil(t, resultErr)
			}

			if !outRandomPath {
				assert.Equal(t, outPath, resultPath)
			}
			assert.Equal(t, outDelete, resultDelete)
		})
	}
}
