package hue

import "testing"

func TestSblackBgf(t *testing.T) {
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
