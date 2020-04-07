package jobs

// GetVariable is used to get the Value for any given Variable for an Action
// returns the value and if the value was actually found
func (stage *Stage) GetVariable(key string, env Environment) (string, bool) {
	if value, found := env[key]; found {
		return value, true
	}

	if value, found := stage.Variables[key]; found {
		return value, true
	}

	return "", false
}
