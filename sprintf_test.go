package hue

import "testing"

func TestSblackf(t *testing.T) {
	str := Sblackf("format %s %s", "test", "word")
	if str != BlackStart+"format test word"+BlackEnd {
		t.Error("Sblackf failed")
	}
}

func TestSredf(t *testing.T) {
	str := Sredf("format %s %s", "test", "word")
	if str != RedStart+"format test word"+RedEnd {
		t.Error("Sredf failed")
	}
}

func TestSgreenf(t *testing.T) {
	str := Sgreenf("format %s %s", "test", "word")
	if str != GreenStart+"format test word"+GreenEnd {
		t.Error("Sgreenf failed")
	}
}

func TestSyellowf(t *testing.T) {
	str := Syellowf("format %s %s", "test", "word")
	if str != YellowStart+"format test word"+YellowEnd {
		t.Error("Syellowf failed")
	}
}

func TestSbluef(t *testing.T) {
	str := Sbluef("format %s %s", "test", "word")
	if str != BlueStart+"format test word"+BlueEnd {
		t.Error("Sbluef failed")
	}
}

func TestSmagentaf(t *testing.T) {
	str := Smagentaf("format %s %s", "test", "word")
	if str != MagentaStart+"format test word"+MagentaEnd {
		t.Error("Smagentaf failed")
	}
}

func TestStealf(t *testing.T) {
	str := Stealf("format %s %s", "test", "word")
	if str != TealStart+"format test word"+TealEnd {
		t.Error("Stealf failed")
	}
}

func TestSwhitef(t *testing.T) {
	str := Swhitef("format %s %s", "test", "word")
	if str != WhiteStart+"format test word"+WhiteEnd {
		t.Error("Swhitef failed")
	}
}

func TestSblackBgf(t *testing.T) {
	str := SblackBgf("format %s %s", "test", "word")
	if str != BlackBgStart+"format test word"+BlackBgEnd {
		t.Error("SblackBgf failed")
	}
}

func TestSblueBgf(t *testing.T) {
	str := SblueBgf("format %s %s", "test", "word")
	if str != BlueBgStart+"format test word"+BlueBgEnd {
		t.Error("SblueBgf failed")
	}
}

func TestSredBgf(t *testing.T) {
	str := SredBgf("format %s %s", "test", "word")
	if str != RedBgStart+"format test word"+RedBgEnd {
		t.Error("SredBgf failed")
	}
}

func TestSgreenBgf(t *testing.T) {
	str := SgreenBgf("format %s %s", "test", "word")
	if str != GreenBgStart+"format test word"+GreenBgEnd {
		t.Error("SgreenBgf failed")
	}
}

func TestSyellowBgf(t *testing.T) {
	str := SyellowBgf("format %s %s", "test", "word")
	if str != YellowBgStart+"format test word"+YellowBgEnd {
		t.Error("SyellowBgf failed")
	}
}

func TestSmagentaBgf(t *testing.T) {
	str := SmagentaBgf("format %s %s", "test", "word")
	if str != MagentaBgStart+"format test word"+MagentaBgEnd {
		t.Error("SmagentaBgf failed")
	}
}

func TestStealBgf(t *testing.T) {
	str := StealBgf("format %s %s", "test", "word")
	if str != TealBgStart+"format test word"+TealBgEnd {
		t.Error("StealBgf failed")
	}
}

func TestSwhiteBgf(t *testing.T) {
	str := SwhiteBgf("format %s %s", "test", "word")
	if str != WhiteBgStart+"format test word"+WhiteBgEnd {
		t.Error("SwhiteBgf failed")
	}
}
