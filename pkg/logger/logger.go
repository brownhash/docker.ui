package logger

import (
	"log"
	"fmt"
)

const (
	SuccessColor = "\033[1;32m%s\033[0m"
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

// Success - log success messages
func Success(message string) {
	formatter := fmt.Sprintf(SuccessColor, message)
	log.Println(formatter)
}

// Info - log informative messages
func Info(message string) {
	formatter := fmt.Sprintf(InfoColor, message)
	log.Println(formatter)
}

// Warn - log warning messages
func Warn(message string) {
	formatter := fmt.Sprintf(WarningColor, message)
	log.Println(formatter)
}

// Error - log error messages
func Error(message string) {
	formatter := fmt.Sprintf(ErrorColor, message)
	log.Fatal(formatter)
}

// Debug - log debugging messages
func Debug(message string) {
	formatter := fmt.Sprintf(DebugColor, message)
	log.Fatal(formatter)
}
