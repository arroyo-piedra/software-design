package adapter

import "fmt"

// ThirdPartyLogger is an existing third-party logging system
type ThirdPartyLogger struct{}

func (l *ThirdPartyLogger) WriteLog(level string, msg string) {
	fmt.Printf("[%s] %s\n", level, msg)
}
