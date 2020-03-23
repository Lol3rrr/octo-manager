package jobs

func (stage *Stage) GetVariable(key string, env Environment) (string, bool) {
	if value, found := env[key]; found {
		return value, true
	}

	if value, found := stage.Variables[key]; found {
		return value, true
	}

	return "", false
}
