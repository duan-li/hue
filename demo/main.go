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

	// demo sprintf.go
	fmt.Println(hue.Sblackf("this is %s %s", "black", "text"))
	fmt.Println(hue.Sredf("this is %s %s", "red", "text"))
	fmt.Println(hue.Sgreenf("this is %s %s", "green", "text"))
	fmt.Println(hue.Syellowf("this is %s %s", "yellow", "text"))
	fmt.Println(hue.Sbluef("this is %s %s", "blue", "text"))
	fmt.Println(hue.Smagentaf("this is %s %s", "magenta", "text"))
	fmt.Println(hue.Stealf("this is %s %s", "teal", "text"))
	fmt.Println(hue.Swhitef("this is %s %s", "white", "text"))
	fmt.Println(hue.SblackBgf("this is %s %s", "black", "background"))
	fmt.Println(hue.SredBgf("this is %s %s", "red", "background"))
	fmt.Println(hue.SgreenBgf("this is %s %s", "green", "background"))
	fmt.Println(hue.SyellowBgf("this is %s %s", "yellow", "background"))
	fmt.Println(hue.SblueBgf("this is %s %s", "blue", "background"))
	fmt.Println(hue.SmagentaBgf("this is %s %s", "magenta", "background"))
	fmt.Println(hue.StealBgf("this is %s %s", "teal", "background"))
	fmt.Println(hue.SwhiteBgf("this is %s %s", "white", "background"))

	// demo println.go
	hue.Blackln("this", "is", "black", "line")
	hue.Redln("this", "is", "red", "line")
	hue.Greenln("this", "is", "green", "line")
	hue.Yellowln("this", "is", "yellow", "line")
	hue.Blueln("this", "is", "blue", "line")
	hue.Magentaln("this", "is", "magenta", "line")
	hue.Tealln("this", "is", "teal", "line")
	hue.Whiteln("this", "is", "white", "line")
	hue.BlackBgln("this", "is", "black", "background", "line")
	hue.RedBgln("this", "is", "red", "background", "line")
	hue.GreenBgln("this", "is", "green", "background", "line")
	hue.YellowBgln("this", "is", "yellow", "background", "line")
	hue.BlueBgln("this", "is", "blue", "background", "line")
	hue.MagentaBgln("this", "is", "magenta", "background", "line")
	hue.TealBgln("this", "is", "teal", "background", "line")
	hue.WhiteBgln("this", "is", "white", "background", "line")

	// Demo cli.go
	hue.Info("this", "is", "info")
	hue.Warn("this", "is", "warn")
	hue.Error("this", "is", "error")
	hue.Success("this", "is", "success")
	hue.Fatal("this", "is", "fatal")

	//// Demo cli.go
	//hue.Blackln("hue.Blackln(\"message\")")
	//hue.Redln("hue.Redln(\"message\")")
	//hue.Greenln("hue.Greenln(\"message\")")
	//hue.Yellowln("hue.Yellowln(\"message\")")
	//hue.Blueln("hue.Blueln(\"message\")")
	//hue.Magentaln("hue.Magentaln(\"message\")")
	//hue.Tealln("hue.Tealln(\"message\")")
	//hue.Whiteln("hue.Whiteln(\"message\")")
	//
	//hue.BlackBgln("hue.BlackBgln(\"message\")")
	//hue.RedBgln("hue.RedBgln(\"message\")")
	//hue.GreenBgln("hue.GreenBgln(\"message\")")
	//hue.YellowBgln("hue.YellowBgln(\"message\")")
	//hue.BlueBgln("hue.BlueBgln(\"message\")")
	//hue.MagentaBgln("hue.MagentaBgln(\"message\")")
	//hue.TealBgln("hue.TealBgln(\"message\")")
	//hue.WhiteBgln("hue.WhiteBgln(\"message\")")
	//
	////
	//hue.Info("hue.Info(\"message\")")
	//hue.Warn("hue.Warn(\"message\")")
	//hue.Error("hue.Error(\"message\")")
	//hue.Fatal("hue.Fatal(\"message\")", false)
	//hue.Success("hue.Success(\"message\")", false)
	//
	//hue.InfoBg("hue.InfoBg(\"message\")")
	//hue.WarnBg("hue.WarnBg(\"message\")")
	//hue.ErrorBg("hue.ErrorBg(\"message\")")
	//hue.FatalBg("hue.FatalBg(\"message\")", false)
	//hue.SuccessBg("hue.SuccessBg(\"message\")", false)
}
