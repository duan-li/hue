package hue

import "fmt"

func Sblackf(s string) string {
	return fmt.Sprintf(Black, s)
}

func Sredf(s string) string {
	return fmt.Sprintf(Red, s)
}

func Sgreenf(s string) string {
	return fmt.Sprintf(Green, s)
}

func Syellowf(s string) string {
	return fmt.Sprintf(Yellow, s)
}

func Sbluef(s string) string {
	return fmt.Sprintf(Blue, s)
}

func Smagentaf(s string) string {
	return fmt.Sprintf(Magenta, s)
}

func Stealf(s string) string {
	return fmt.Sprintf(Teal, s)
}

func Swhitef(s string) string {
	return fmt.Sprintf(White, s)
}

func SblackBgf(s string) string {
	return fmt.Sprintf(BlackBg, s)
}

func SredBgf(s string) string {
	return fmt.Sprintf(RedBg, s)
}

func SgreenBgf(s string) string {
	return fmt.Sprintf(GreenBg, s)
}

func SyellowBgf(s string) string {
	return fmt.Sprintf(YellowBg, s)
}

func SblueBgf(s string) string {
	return fmt.Sprintf(BlueBg, s)
}

func SmagentaBgf(s string) string {
	return fmt.Sprintf(MagentaBg, s)
}

func StealBgf(s string) string {
	return fmt.Sprintf(TealBg, s)
}

func SwhiteBgf(s string) string {
	return fmt.Sprintf(WhiteBg, s)
}
