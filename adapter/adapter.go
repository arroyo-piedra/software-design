package adapter

// Logger is the application's expected interface
type Logger interface {
	Log(message string)
}

// LoggerAdapter adapts ThirdPartyLogger to the Logger interface
type LoggerAdapter struct {
	thirdPartyLogger *ThirdPartyLogger
}

// Log adapts the interface to use the third-party method
func (a *LoggerAdapter) Log(message string) {
	a.thirdPartyLogger.WriteLog("INFO", message) // Maps to INFO level
}

// This is a struct to be invoked by main.go
type AdapterPattern struct{}

func (a AdapterPattern) ProcessPattern() bool {
	// Create an instance of the third-party logger
	thirdPartyLogger := &ThirdPartyLogger{}

	// Use the adapter to bridge the interfaces
	adapter := &LoggerAdapter{thirdPartyLogger: thirdPartyLogger}

	// Client code interacts with the expected interface
	adapter.Log("Adapter pattern in action!") // Output: [INFO] Adapter pattern in action!
	return true
}
