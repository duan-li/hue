package main

import (
	"fmt"
	"github.com/duan-li/hue"
)

func main() {
	// Demo sprint.go
	fmt.Println(hue.Sblack("this", "is", "black", "text", "by", "Sblack"))
	fmt.Println(hue.Sred("this", "is", "red", "text", "by", "Sred"))
	fmt.Println(hue.Sgreen("this", "is", "green", "text", "by", "Sgreen"))
	fmt.Println(hue.Syellow("this", "is", "yellow", "text", "by", "Syellow"))
	fmt.Println(hue.Sblue("this", "is", "blue", "text", "by", "Sblue"))
	fmt.Println(hue.Smagenta("this", "is", "magenta", "text", "by", "Smagenta"))
	fmt.Println(hue.Steal("this", "is", "teal", "text", "by", "Steal"))
	fmt.Println(hue.Swhite("this", "is", "white", "text", "by", "Swhite"))
	fmt.Println(hue.SblackBg("this", "is", "black", "background", "by", "SblackBg"))
	fmt.Println(hue.SredBg("this", "is", "red", "background", "by", "SredBg"))
	fmt.Println(hue.SgreenBg("this", "is", "green", "background", "by", "SgreenBg"))
	fmt.Println(hue.SyellowBg("this", "is", "yellow", "background", "by", "SyellowBg"))
	fmt.Println(hue.SblueBg("this", "is", "blue", "background", "by", "SblueBg"))
	fmt.Println(hue.SmagentaBg("this", "is", "magenta", "background", "by", "SmagentaBg"))
	fmt.Println(hue.StealBg("this", "is", "teal", "background", "by", "StealBg"))
	fmt.Println(hue.SwhiteBg("this", "is", "white", "background", "by", "SwhiteBg"))

	// demo sprintf.go
	fmt.Println(hue.Sblackf("this is %s %s %s %s", "black", "text", "by", "Sblackf"))
	fmt.Println(hue.Sredf("this is %s %s %s %s", "red", "text", "by", "Sredf"))
	fmt.Println(hue.Sgreenf("this is %s %s %s %s", "green", "text", "by", "Sgreenf"))
	fmt.Println(hue.Syellowf("this is %s %s %s %s", "yellow", "text", "by", "Syellowf"))
	fmt.Println(hue.Sbluef("this is %s %s %s %s", "blue", "text", "by", "Sbluef"))
	fmt.Println(hue.Smagentaf("this is %s %s %s %s", "magenta", "text", "by", "Smagentaf"))
	fmt.Println(hue.Stealf("this is %s %s %s %s", "teal", "text", "by", "Stealf"))
	fmt.Println(hue.Swhitef("this is %s %s %s %s", "white", "text", "by", "Swhitef"))
	fmt.Println(hue.SblackBgf("this is %s %s %s %s", "black", "background", "by", "SblackBgf"))
	fmt.Println(hue.SredBgf("this is %s %s %s %s", "red", "background", "by", "SredBgf"))
	fmt.Println(hue.SgreenBgf("this is %s %s %s %s", "green", "background", "by", "SgreenBgf"))
	fmt.Println(hue.SyellowBgf("this is %s %s %s %s", "yellow", "background", "by", "SyellowBgf"))
	fmt.Println(hue.SblueBgf("this is %s %s %s %s", "blue", "background", "by", "SblueBgf"))
	fmt.Println(hue.SmagentaBgf("this is %s %s %s %s", "magenta", "background", "by", "SmagentaBgf"))
	fmt.Println(hue.StealBgf("this is %s %s %s %s", "teal", "background", "by", "StealBgf"))
	fmt.Println(hue.SwhiteBgf("this is %s %s %s %s", "white", "background", "by", "SwhiteBgf"))

	// demo println.go
	hue.Blackln("this", "is", "black", "line", "by", "Blackln")
	hue.Redln("this", "is", "red", "line", "by", "Redln")
	hue.Greenln("this", "is", "green", "line", "by", "Greenln")
	hue.Yellowln("this", "is", "yellow", "line", "by", "Yellowln")
	hue.Blueln("this", "is", "blue", "line", "by", "Blueln")
	hue.Magentaln("this", "is", "magenta", "line", "by", "Magentaln")
	hue.Tealln("this", "is", "teal", "line", "by", "Tealln")
	hue.Whiteln("this", "is", "white", "line", "by", "Whiteln")
	hue.BlackBgln("this", "is", "black", "background", "line", "by", "BlackBgln")
	hue.RedBgln("this", "is", "red", "background", "line", "by", "RedBgln")
	hue.GreenBgln("this", "is", "green", "background", "line", "by", "GreenBgln")
	hue.YellowBgln("this", "is", "yellow", "background", "line", "by", "YellowBgln")
	hue.BlueBgln("this", "is", "blue", "background", "line", "by", "BlueBgln")
	hue.MagentaBgln("this", "is", "magenta", "background", "line", "by", "MagentaBgln")
	hue.TealBgln("this", "is", "teal", "background", "line", "by", "TealBgln")
	hue.WhiteBgln("this", "is", "white", "background", "line", "by", "WhiteBgln")

	// Demo cli.go
	hue.Info("this", "is", "info")
	hue.Warn("this", "is", "warn")
	hue.Error("this", "is", "error")
	hue.Success("this", "is", "success")
	hue.Fatal("this", "is", "fatal")
}
