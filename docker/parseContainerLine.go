package docker

import (
  "errors"
  "strings"
)

func parseContainerLine(line string) (Container, error) {
  rawParts := strings.Split(line, "  ")
  parts := make([]string, 0)

  for _, rawPart := range rawParts {
    if len(rawPart) <= 0 {
      continue
    }

    if rawPart[0] == byte(' ') {
      rawPart = rawPart[1:]
    }

    parts = append(parts, rawPart)
  }

  if len(parts) < 7 {
    return Container{}, errors.New("The Line is missing data")
  }

  return Container{
    ID: parts[0],
    Image: parts[1],
    Command: parts[2],
    Created: parts[3],
    Status: parts[4],
    Ports: parts[5],
    Names: parts[6],
  }, nil
}
