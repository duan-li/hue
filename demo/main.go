package main

import (
	"fmt"
	"github.com/duan-li/hue"
)

func main() {
	// Demo sprint.go
	fmt.Println(hue.Sblack("this", "is", "black", "text"))
	fmt.Println(hue.Sred("this", "is", "red", "text"))
	fmt.Println(hue.Sgreen("this", "is", "green", "text"))
	fmt.Println(hue.Syellow("this", "is", "yellow", "text"))
	fmt.Println(hue.Sblue("this", "is", "blue", "text"))
	fmt.Println(hue.Smagenta("this", "is", "magenta", "text"))
	fmt.Println(hue.Steal("this", "is", "teal", "text"))
	fmt.Println(hue.Swhite("this", "is", "white", "text"))
	fmt.Println(hue.SblackBg("this", "is", "black", "background"))
	fmt.Println(hue.SredBg("this", "is", "red", "background"))
	fmt.Println(hue.SgreenBg("this", "is", "green", "background"))
	fmt.Println(hue.SyellowBg("this", "is", "yellow", "background"))
	fmt.Println(hue.SblueBg("this", "is", "blue", "background"))
	fmt.Println(hue.SmagentaBg("this", "is", "magenta", "background"))
	fmt.Println(hue.StealBg("this", "is", "teal", "background"))
	fmt.Println(hue.SwhiteBg("this", "is", "white", "background"))

	// Demo cli.go
	hue.Blackln("hue.Blackln(\"message\")")
	hue.Redln("hue.Redln(\"message\")")
	hue.Greenln("hue.Greenln(\"message\")")
	hue.Yellowln("hue.Yellowln(\"message\")")
	hue.Blueln("hue.Blueln(\"message\")")
	hue.Magentaln("hue.Magentaln(\"message\")")
	hue.Tealln("hue.Tealln(\"message\")")
	hue.Whiteln("hue.Whiteln(\"message\")")

	hue.BlackBgln("hue.BlackBgln(\"message\")")
	hue.RedBgln("hue.RedBgln(\"message\")")
	hue.GreenBgln("hue.GreenBgln(\"message\")")
	hue.YellowBgln("hue.YellowBgln(\"message\")")
	hue.BlueBgln("hue.BlueBgln(\"message\")")
	hue.MagentaBgln("hue.MagentaBgln(\"message\")")
	hue.TealBgln("hue.TealBgln(\"message\")")
	hue.WhiteBgln("hue.WhiteBgln(\"message\")")

	//
	hue.Info("hue.Info(\"message\")")
	hue.Warn("hue.Warn(\"message\")")
	hue.Error("hue.Error(\"message\")")
	hue.Fatal("hue.Fatal(\"message\")", false)
	hue.Success("hue.Success(\"message\")", false)

	hue.InfoBg("hue.InfoBg(\"message\")")
	hue.WarnBg("hue.WarnBg(\"message\")")
	hue.ErrorBg("hue.ErrorBg(\"message\")")
	hue.FatalBg("hue.FatalBg(\"message\")", false)
	hue.SuccessBg("hue.SuccessBg(\"message\")", false)
}
