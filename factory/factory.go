package factory

import "fmt"

const GoLanguage = "go"
const PythonLanguage = "python"

// This is an interface that defines the methods that the concrete structs will implement
type ProcessFileForLanguage interface {
	// Reads a file based on the language
	ReadFile() string
	// Processes a file based on the language
	ProcessFile() string
}

// Creates the requested struct based on the language
func ProcessFileForLanguageFactory(language string) (ProcessFileForLanguage, error) {
	switch language {
	case GoLanguage:
		return GoFile{}, nil
	case PythonLanguage:
		return PythonFile{}, nil
	default:
		return nil, fmt.Errorf("language not supported: %s", language)
	}
}

// This is a struct to be invoked by main.go
type FactoryPattern struct{}

func (f FactoryPattern) ProcessPattern() bool {
	fmt.Println("------------ Factory Pattern ------------")

	fmt.Println("Factory Pattern: Creating Go Factory")
	goFactory, err := ProcessFileForLanguageFactory(GoLanguage)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Calling Go Factory methods")
	fmt.Println(goFactory.ReadFile())
	fmt.Println(goFactory.ProcessFile())

	fmt.Println("Factory Pattern: Creating Python Factory")
	pythonFactory, err := ProcessFileForLanguageFactory(PythonLanguage)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Calling Python Factory methods")
	fmt.Println(pythonFactory.ReadFile())
	fmt.Println(pythonFactory.ProcessFile())
	return true
}
