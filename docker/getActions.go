package docker

// GetActions returns all Actions of the Docker-Module
func (m *Module) GetActions() []string {
	return []string{
		"pull",
		"compose up",
	}
}