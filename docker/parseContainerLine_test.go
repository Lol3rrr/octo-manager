package docker

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestParseContainerLine(t *testing.T) {
  tables := []struct{
    Name string
    InputLine string
    ResultContainer Container
    ResultError bool
  }{
    {
      Name: "Valid Input",
      InputLine: "testID  testImage  testCommand  testCreated  testStatus  testPorts  testNames",
      ResultContainer: Container{
        ID: "testID",
        Image: "testImage",
        Command: "testCommand",
        Created: "testCreated",
        Status: "testStatus",
        Ports: "testPorts",
        Names: "testNames",
      },
      ResultError: false,
    },
    {
      Name: "Too few elements in the line",
      InputLine: "testID  testImage  testCommand  testCreated  testStatus  testPorts",
      ResultContainer: Container{},
      ResultError: true,
    },
  }

  for _, table := range tables {
    inLine := table.InputLine
    resultContainer := table.ResultContainer
    resultError := table.ResultError

    t.Run(table.Name, func(t *testing.T) {
      outputContainer, outputError := parseContainerLine(inLine)

      assert.Equal(t, resultContainer, outputContainer)
      if resultError {
        assert.NotNil(t, outputError)
      } else {
        assert.Nil(t, outputError)
      }
    })
  }
}
