package hue

import "fmt"

func Sblack(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = BlackStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlackEnd
	return fmt.Sprint(an...)
}

func Sred(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = RedStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = RedEnd
	return fmt.Sprint(an...)
}

func Sgreen(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = GreenStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = GreenEnd
	return fmt.Sprint(an...)
}

func Syellow(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = YellowStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = YellowEnd
	return fmt.Sprint(an...)
}

func Sblue(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = BlueStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlueEnd
	return fmt.Sprint(an...)
}

func Smagenta(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = MagentaStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = MagentaEnd
	return fmt.Sprint(an...)
}

func Steal(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = TealStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = TealEnd
	return fmt.Sprint(an...)
}

func Swhite(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = WhiteStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = WhiteEnd
	return fmt.Sprint(an...)
}

func SblackBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = BlackBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlackBgEnd
	return fmt.Sprint(an...)
}

func SredBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = RedBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = RedBgEnd
	return fmt.Sprint(an...)
}

func SgreenBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = GreenBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = GreenBgEnd
	return fmt.Sprint(an...)
}

func SyellowBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = YellowBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = YellowBgEnd
	return fmt.Sprint(an...)
}

func SblueBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = BlueBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = BlueBgEnd
	return fmt.Sprint(an...)
}

func SmagentaBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = MagentaBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = MagentaBgEnd
	return fmt.Sprint(an...)
}

func StealBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = TealBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = TealBgEnd
	return fmt.Sprint(an...)
}

func SwhiteBg(s ...string) string {
	an := make([]interface{}, len(s)+2)
	an[0] = WhiteBgStart
	for i, v := range s {
		i++
		an[i] = v
	}
	an[len(s)+1] = WhiteBgEnd
	return fmt.Sprint(an...)
}
