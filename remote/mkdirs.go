package remote

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (s *session) mkdirs(dir string) error {
	dirs := strings.Split(dir, "/")
	for i := range dirs {
		tmpPath := filepath.Join(dirs[:i+1]...)

		cdCmd := fmt.Sprintf("cd %s", tmpPath)
		_, err := s.Command(cdCmd)
		if err != nil {
			mkdirCmd := fmt.Sprintf("mkdir %s", tmpPath)
			_, err = s.Command(mkdirCmd)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
