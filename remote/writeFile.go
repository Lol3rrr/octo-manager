package remote

import "fmt"

func (s *session) WriteFile(path, content string) error {
	rawWriteCmd := fmt.Sprintf("cat > %s \n", path)

	writeCmd := rawWriteCmd + content
	_, err := s.ShellCommand(writeCmd)
	if err != nil {
		return err
	}

	return nil
}
