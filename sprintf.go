package hue

import "fmt"

// Sblackf returns a string with black color
func Sblackf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Black, s)
}

// Sredf returns a string with red color
func Sredf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Red, s)
}

// Sgreenf returns a string with green color
func Sgreenf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Green, s)
}

// Syellowf returns a string with yellow color
func Syellowf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Yellow, s)
}

// Sbluef returns a string with blue color
func Sbluef(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Blue, s)
}

// Smagentaf returns a string with magenta color
func Smagentaf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Magenta, s)
}

// Stealf returns a string with the color
func Stealf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(Teal, s)
}

// Swhitef writes a string with black color
func Swhitef(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(White, s)
}

// SblackBgf returns a string with black background color
func SblackBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(BlackBg, s)
}

// SredBgf returns a string with red background color
func SredBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(RedBg, s)
}

// SgreenBgf returns a string with green background color
func SgreenBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(GreenBg, s)
}

// SyellowBgf returns a string with yellow background color
func SyellowBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(YellowBg, s)
}

// SblueBgf returns a string with blue background color
func SblueBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(BlueBg, s)
}

// SmagentaBgf returns a string with magenta background color
func SmagentaBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(MagentaBg, s)
}

// StealBgf returns a string with teal background color
func StealBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(TealBg, s)
}

// SwhiteBgf returns a string with white background color
func SwhiteBgf(format string, a ...any) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf(WhiteBg, s)
}
