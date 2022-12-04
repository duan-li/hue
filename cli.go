package hue

import "os"

func Info(message string) {
	Blueln(message)
}

func Warn(message string) {
	Yellowln(message)
}

func Error(message string) {
	Redln(message)
}

func Fatal(message string, exit bool) {
	Magentaln(message)
	if exit {
		os.Exit(1)
	}
}

func Success(message string, exit bool) {
	Greenln(message)
	if exit {
		os.Exit(0)
	}
}

func InfoBg(message string) {
	BlueBgln(message)
}

func WarnBg(message string) {
	YellowBgln(message)
}

func ErrorBg(message string) {
	RedBgln(message)
}

func FatalBg(message string, exit bool) {
	MagentaBgln(message)
	if exit {
		os.Exit(1)
	}
}

func SuccessBg(message string, exit bool) {
	GreenBgln(message)
	if exit {
		os.Exit(0)
	}
}
