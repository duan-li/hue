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
	an := make([]interface{}, len(s)+2)
	an[0] = RedStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = RedEnd
	return fmt.Println(an...)
}

// Greenln prints a message in green
func Greenln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = GreenStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = GreenEnd
	return fmt.Println(an...)
}

// Yellowln prints a message in yellow
func Yellowln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = YellowStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = YellowEnd
	return fmt.Println(an...)
}

// Blueln prints a message in blue
func Blueln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = BlueStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlueEnd
	return fmt.Println(an...)
}

// Magentaln prints a message in magenta
func Magentaln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = MagentaStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = MagentaEnd
	return fmt.Println(an...)
}

// Tealln prints a message in teal
func Tealln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = TealStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = TealEnd
	return fmt.Println(an...)
}

// Whiteln prints a message in white
func Whiteln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = WhiteStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = WhiteEnd
	return fmt.Println(an...)
}

// BlackBgln prints a message in black background
func BlackBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = BlackBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlackBgEnd
	return fmt.Println(an...)
}

// RedBgln prints a message in red background
func RedBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = RedBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = RedBgEnd
	return fmt.Println(an...)
}

// GreenBgln prints a message in green background
func GreenBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = GreenBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = GreenBgEnd
	return fmt.Println(an...)
}

// YellowBgln prints a message in yellow background
func YellowBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = YellowBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = YellowBgEnd
	return fmt.Println(an...)
}

// BlueBgln prints a message in blue background
func BlueBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = BlueBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlueBgEnd
	return fmt.Println(an...)
}

// MagentaBgln prints a message in magenta background
func MagentaBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = MagentaBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = MagentaBgEnd
	return fmt.Println(an...)
}

// TealBgln prints a message in teal background
func TealBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = TealBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = TealBgEnd
	return fmt.Println(an...)
}

// WhiteBgln prints a message in white background
func WhiteBgln(s ...string) (n int, err error) {
	an := make([]interface{}, len(s)+2)
	an[0] = WhiteBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = WhiteBgEnd
	return fmt.Println(an...)
}
