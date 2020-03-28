package remote

func isDir(path string) bool {
	if len(path) <= 0 {
		return false
	}

	return path[len(path)-1] == '/'
}
