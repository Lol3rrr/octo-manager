package remote

// Session is a simple interface that represents a
// connection to a server over ssh
type Session interface {
	// Command simply executes a command on the Server and returns the result
	Command(cmd string) (string, error)
}
