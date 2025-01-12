package main

import fmt "fmt"

// package level variable - can be accessed in app.go
// variable 'logLevel' is set to 0 in app.go file & this file uses that variable

func logInfo(msg string) {
	if logLevel == 0 {
		fmt.Printf("[%s] %s\n", getLogLevel(), msg)
	}
}

func getLogLevel() string {
	if logLevel == 0 {
		return "INFO"
	} else {
		return "TRACE"
	}
}
