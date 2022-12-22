package hue

import "testing"

func TestSblack(t *testing.T) {
	if Sblack("test", "word") != BlackStart+"test"+"word"+BlackEnd {
		t.Error("Sblack failed")
	}
}

func TestSred(t *testing.T) {
	if Sred("test", "word") != RedStart+"test"+"word"+RedEnd {
		t.Error("Sred failed")
	}
}

func TestSgreen(t *testing.T) {
	if Sgreen("test", "word") != GreenStart+"test"+"word"+GreenEnd {
		t.Error("Sgreen failed")
	}
}

func TestSyellow(t *testing.T) {
	if Syellow("test", "word") != YellowStart+"test"+"word"+YellowEnd {
		t.Error("Syellow failed")
	}
}

func TestSblue(t *testing.T) {
	if Sblue("test", "word") != BlueStart+"test"+"word"+BlueEnd {
		t.Error("Sblue failed")
	}
}

func TestSmagenta(t *testing.T) {
	if Smagenta("test", "word") != MagentaStart+"test"+"word"+MagentaEnd {
		t.Error("Smagenta failed")
	}
}

func TestSteal(t *testing.T) {
	if Steal("test", "word") != TealStart+"test"+"word"+TealEnd {
		t.Error("Steal failed")
	}
}

func TestSwhite(t *testing.T) {
	if Swhite("test", "word") != WhiteStart+"test"+"word"+WhiteEnd {
		t.Error("Swhite failed")
	}
}

func TestSblackbg(t *testing.T) {
	if SblackBg("test", "word") != BlackBgStart+"test"+"word"+BlackBgEnd {
		t.Error("SblackBg failed")
	}
}

func TestSredbg(t *testing.T) {
	if SredBg("test", "word") != RedBgStart+"test"+"word"+RedBgEnd {
		t.Error("SredBg failed")
	}
}

func TestSgreenbg(t *testing.T) {
	if SgreenBg("test", "word") != GreenBgStart+"test"+"word"+GreenBgEnd {
		t.Error("SgreenBg failed")
	}
}

func TestSyellowbg(t *testing.T) {
	if SyellowBg("test", "word") != YellowBgStart+"test"+"word"+YellowBgEnd {
		t.Error("SyellowBg failed")
	}
}

func TestSbluebg(t *testing.T) {
	if SblueBg("test", "word") != BlueBgStart+"test"+"word"+BlueBgEnd {
		t.Error("SblueBg failed")
	}
}

func TestSmagentabg(t *testing.T) {
	if SmagentaBg("test", "word") != MagentaBgStart+"test"+"word"+MagentaBgEnd {
		t.Error("SmagentaBg failed")
	}
}

func TestStealbg(t *testing.T) {
	if StealBg("test", "word") != TealBgStart+"test"+"word"+TealBgEnd {
		t.Error("StealBg failed")
	}
}

func TestSwhitebg(t *testing.T) {
	if SwhiteBg("test", "word") != WhiteBgStart+"test"+"word"+WhiteBgEnd {
		t.Error("SwhiteBg failed")
	}
}
