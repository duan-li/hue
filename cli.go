package hue

import "os"

// Info prints a message in blue
func Info(message string) {
	Blueln(message)
}

// Warn prints a message in yellow
func Warn(message string) {
	Yellowln(message)
}

// Error prints a message in red
func Error(message string) {
	Redln(message)
}

// Fatal prints a message in magenta and exit with code 1
func Fatal(message string, exit bool) {
	Magentaln(message)
	if exit {
		os.Exit(1)
	}
}

// Success prints a message in green and exit with code 0
func Success(message string, exit bool) {
	Greenln(message)
	if exit {
		os.Exit(0)
	}
}

// InfoBg prints a message in blue background
func InfoBg(message string) {
	BlueBgln(message)
}

// WarnBg prints a message in yellow background
func WarnBg(message string) {
	YellowBgln(message)
}

// ErrorBg prints a message in red background
func ErrorBg(message string) {
	RedBgln(message)
}

// FatalBg prints a message in magenta background and exit with code 1
func FatalBg(message string, exit bool) {
	MagentaBgln(message)
	if exit {
		os.Exit(1)
	}
}

// SuccessBg prints a message in green background and exit with code 0
func SuccessBg(message string, exit bool) {
	GreenBgln(message)
	if exit {
		os.Exit(0)
	}
}
