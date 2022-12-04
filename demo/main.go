package main

import "github.com/duan-li/hue"

func main() {
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
