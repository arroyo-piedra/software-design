package factory

// Example of a go implementation that reads and processes .out files
type GoFile struct{}

func (g GoFile) ReadFile() string {
	return "Go reads .out files"
}

func (g GoFile) ProcessFile() string {
	return "Go processes .out files"
}
