package docker

// Container represents a single Container instance
type Container struct {
	ID      string
	Image   string
	Command string
	Created string
	Status  string
	Ports   string
	Names   string
}

// Module is simply an object that all the needed functions will be tied to
type Module struct {
}
