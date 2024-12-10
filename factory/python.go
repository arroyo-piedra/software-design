package factory

// Example of a python implementation that reads and processes .out files
type PythonFile struct{}

func (g PythonFile) ReadFile() string {
	return "Python reads .txt files"
}

func (g PythonFile) ProcessFile() string {
	return "Python processes .txt files"
}
