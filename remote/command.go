package remote

func (s *session) Command(cmd string) (string, error) {
	rawOut, err := s.SSHClient.Cmd(cmd).Output()

	return string(rawOut), err
}
