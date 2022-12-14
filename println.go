package hue

import "fmt"

// Blackln prints a message in black
func Blackln(s string) {
	fmt.Println(Sblackf(s))
}

// Redln prints a message in red
func Redln(s string) {
	fmt.Println(Sredf(s))
}

// Greenln prints a message in green
func Greenln(s string) {
	fmt.Println(Sgreenf(s))
}

// Yellowln prints a message in yellow
func Yellowln(s string) {
	fmt.Println(Syellowf(s))
}

// Blueln prints a message in blue
func Blueln(s string) {
	fmt.Println(Sbluef(s))
}

// Magentaln prints a message in magenta
func Magentaln(s string) {
	fmt.Println(Smagentaf(s))
}

// Tealln prints a message in teal
func Tealln(s string) {
	fmt.Println(Stealf(s))
}

// Whiteln prints a message in white
func Whiteln(s string) {
	fmt.Println(Swhitef(s))
}

// BlackBgln prints a message in black background
func BlackBgln(s string) {
	fmt.Println(SblackBgf(s))
}

// RedBgln prints a message in red background
func RedBgln(s string) {
	fmt.Println(SredBgf(s))
}

// GreenBgln prints a message in green background
func GreenBgln(s string) {
	fmt.Println(SgreenBgf(s))
}

// YellowBgln prints a message in yellow background
func YellowBgln(s string) {
	fmt.Println(SyellowBgf(s))
}

// BlueBgln prints a message in blue background
func BlueBgln(s string) {
	fmt.Println(SblueBgf(s))
}

// MagentaBgln prints a message in magenta background
func MagentaBgln(s string) {
	fmt.Println(SmagentaBgf(s))
}

// TealBgln prints a message in teal background
func TealBgln(s string) {
	fmt.Println(StealBgf(s))
}

// WhiteBgln prints a message in white background
func WhiteBgln(s string) {
	fmt.Println(SwhiteBgf(s))
}
