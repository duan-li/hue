package hue

import "os"

// Info prints a message in blue
func Info(message ...string) {
	Blueln(message...)
}

// Warn prints a message in yellow
func Warn(message ...string) {
	Yellowln(message...)
}

// Error prints a message in red
func Error(message ...string) {
	Redln(message...)
}

// Fatal prints a message in magenta and exit with code 1
func Fatal(message ...string) {
	RedBgln(message...)
	os.Exit(1)
}

// Success prints a message in green and exit with code 0
func Success(message ...string) {
	Greenln(message...)
}
