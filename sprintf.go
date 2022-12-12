package hue

import "fmt"

// Sblackf returns a string with black color
func Sblackf(s string) string {
	return fmt.Sprintf(Black, s)
}

// Sredf returns a string with red color
func Sredf(s string) string {
	return fmt.Sprintf(Red, s)
}

// Sgreenf returns a string with green color
func Sgreenf(s string) string {
	return fmt.Sprintf(Green, s)
}

// Syellowf returns a string with yellow color
func Syellowf(s string) string {
	return fmt.Sprintf(Yellow, s)
}

// Sbluef returns a string with blue color
func Sbluef(s string) string {
	return fmt.Sprintf(Blue, s)
}

// Smagentaf returns a string with magenta color
func Smagentaf(s string) string {
	return fmt.Sprintf(Magenta, s)
}

// Stealf returns a string with the color
func Stealf(s string) string {
	return fmt.Sprintf(Teal, s)
}

// Swhitef writes a string with black color
func Swhitef(s string) string {
	return fmt.Sprintf(White, s)
}

// SblackBgf returns a string with black background color
func SblackBgf(s string) string {
	return fmt.Sprintf(BlackBg, s)
}

// SredBgf returns a string with red background color
func SredBgf(s string) string {
	return fmt.Sprintf(RedBg, s)
}

// SgreenBgf returns a string with green background color
func SgreenBgf(s string) string {
	return fmt.Sprintf(GreenBg, s)
}

// SyellowBgf returns a string with yellow background color
func SyellowBgf(s string) string {
	return fmt.Sprintf(YellowBg, s)
}

// SblueBgf returns a string with blue background color
func SblueBgf(s string) string {
	return fmt.Sprintf(BlueBg, s)
}

// SmagentaBgf returns a string with magenta background color
func SmagentaBgf(s string) string {
	return fmt.Sprintf(MagentaBg, s)
}

// StealBgf returns a string with teal background color
func StealBgf(s string) string {
	return fmt.Sprintf(TealBg, s)
}

// SwhiteBgf returns a string with white background color
func SwhiteBgf(s string) string {
	return fmt.Sprintf(WhiteBg, s)
}
