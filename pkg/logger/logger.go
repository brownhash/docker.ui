package logger

import (
	"os"
	"log"
	"fmt"
	"net/http"
)

const (
	SuccessColor = "\033[1;32m%s\033[0m"
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

// SetLogLevel - set global logging level
func SetLogLevel(logLevel string) {
	os.Setenv("DUI_LOGGING_LEVEL", logLevel)
}

func getLogLevel() string {
	return os.Getenv("DUI_LOGGING_LEVEL")
}

// Debug - log debugging messages
func Debug(message interface{}) {
	if getLogLevel() == "DEBUG" {
		formatter := fmt.Sprintf(DebugColor, message)
		log.Println(formatter)
	}
}

// Debugf - log debugging messages in same line
func Debugf(message interface{}) {
	if getLogLevel() == "DEBUG" {
		formatter := fmt.Sprintf(DebugColor, message)
		log.Printf(formatter)
	}
}

// Info - log informative messages
func Info(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" {
		formatter := fmt.Sprintf(InfoColor, message)
		log.Println(formatter)
	}
}

// Infof - log informative messages in same line
func Infof(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" {
		formatter := fmt.Sprintf(InfoColor, message)
		log.Printf(formatter)
	}
}

// Warn - log warning messages
func Warn(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" || getLogLevel() == "WARN" {
		formatter := fmt.Sprintf(WarningColor, message)
		log.Println(formatter)
	}
}

// Warnf - log warning messages in same line
func Warnf(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" || getLogLevel() == "WARN" {
		formatter := fmt.Sprintf(WarningColor, message)
		log.Printf(formatter)
	}
}

// Success - log success messages
func Success(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" || getLogLevel() == "WARN" {
		formatter := fmt.Sprintf(SuccessColor, message)
		log.Println(formatter)
	}
}

// Successf - log success messages in same line
func Successf(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" || getLogLevel() == "WARN" {
		formatter := fmt.Sprintf(SuccessColor, message)
		log.Printf(formatter)
	}
}

// Error - log error messages
func Error(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" || getLogLevel() == "WARN" || getLogLevel() == "ERROR" {
		formatter := fmt.Sprintf(ErrorColor, message)
		log.Fatal(formatter)
	}
}

// Errorf - log error messages in same line
func Errorf(message interface{}) {
	if getLogLevel() == "DEBUG" || getLogLevel() == "INFO" || getLogLevel() == "WARN" || getLogLevel() == "ERROR" {
		formatter := fmt.Sprintf(ErrorColor, message)
		log.Fatalf(formatter)
	}
}

// LogRequest - log http requests
func LogRequest(handler http.Handler) http.Handler {
	log.SetFlags(log.Ldate | log.Ltime)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formatter := fmt.Sprintf("%s %s", fmt.Sprintf(SuccessColor, r.Method), fmt.Sprintf(DebugColor, r.URL))
		log.Printf(formatter)

		handler.ServeHTTP(w, r)
	})
}
