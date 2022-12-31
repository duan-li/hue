package hue

import (
	"fmt"
	"strings"
)

// Blackln prints a message in black
func Blackln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(BlackStart + str + BlackEnd)
}

// Redln prints a message in red
func Redln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(RedStart + str + RedEnd)
}

// Greenln prints a message in green
func Greenln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(GreenStart + str + GreenEnd)
}

// Yellowln prints a message in yellow
func Yellowln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(YellowStart + str + YellowEnd)
}

// Blueln prints a message in blue
func Blueln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(BlueStart + str + BlueEnd)
}

// Magentaln prints a message in magenta
func Magentaln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(MagentaStart + str + MagentaEnd)
}

// Tealln prints a message in teal
func Tealln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(TealStart + str + TealEnd)
}

// Whiteln prints a message in white
func Whiteln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(WhiteStart + str + WhiteEnd)
}

// BlackBgln prints a message in black background
func BlackBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(BlackBgStart + str + BlackBgEnd)
}

// RedBgln prints a message in red background
func RedBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(RedBgStart + str + RedBgEnd)
}

// GreenBgln prints a message in green background
func GreenBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(GreenBgStart + str + GreenBgEnd)
}

// YellowBgln prints a message in yellow background
func YellowBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(YellowBgStart + str + YellowBgEnd)
}

// BlueBgln prints a message in blue background
func BlueBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(BlueBgStart + str + BlueBgEnd)
}

// MagentaBgln prints a message in magenta background
func MagentaBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(MagentaBgStart + str + MagentaBgEnd)
}

// TealBgln prints a message in teal background
func TealBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(TealBgStart + str + TealBgEnd)
}

// WhiteBgln prints a message in white background
func WhiteBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s))
	for i, v := range s {
		an[i] = v
	}
	str := fmt.Sprintln(an...)
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return fmt.Println(WhiteBgStart + str + WhiteBgEnd)
}
