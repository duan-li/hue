package hue

import (
	"io"
	"os"
	"testing"
)

func TestBlackln(t *testing.T) {
	output := captureOutput(func() {
		Blackln("test", "word")
	})

	if output != BlackStart+"test word"+BlackEnd+"\n" {
		t.Error("Blackln failed")
	}
}

func TestRedln(t *testing.T) {
	output := captureOutput(func() {
		Redln("test", "word")
	})

	if output != RedStart+"test word"+RedEnd+"\n" {
		t.Error("Redln failed")
	}
}

func TestGreenln(t *testing.T) {
	output := captureOutput(func() {
		Greenln("test", "word")
	})

	if output != GreenStart+"test word"+GreenEnd+"\n" {
		t.Error("Greenln failed")
	}
}

func TestYellowln(t *testing.T) {
	output := captureOutput(func() {
		Yellowln("test", "word")
	})

	if output != YellowStart+"test word"+YellowEnd+"\n" {
		t.Error("Yellowln failed")
	}
}

func TestBlueln(t *testing.T) {
	output := captureOutput(func() {
		Blueln("test", "word")
	})

	if output != BlueStart+"test word"+BlueEnd+"\n" {
		t.Error("Blueln failed")
	}
}

func TestMagentaln(t *testing.T) {
	output := captureOutput(func() {
		Magentaln("test", "word")
	})

	if output != MagentaStart+"test word"+MagentaEnd+"\n" {
		t.Error("Magentaln failed")
	}
}

func TestTealln(t *testing.T) {
	output := captureOutput(func() {
		Tealln("test", "word")
	})

	if output != TealStart+"test word"+TealEnd+"\n" {
		t.Error("Tealln failed")
	}
}

func TestWhiteln(t *testing.T) {
	output := captureOutput(func() {
		Whiteln("test", "word")
	})

	if output != WhiteStart+"test word"+WhiteEnd+"\n" {
		t.Error("Whiteln failed")
	}
}

func TestBlackBgln(t *testing.T) {
	output := captureOutput(func() {
		BlackBgln("test", "word")
	})

	if output != BlackBgStart+"test word"+BlackBgEnd+"\n" {
		t.Error("BlackBgln failed")
	}
}

func TestRedBgln(t *testing.T) {
	output := captureOutput(func() {
		RedBgln("test", "word")
	})

	if output != RedBgStart+"test word"+RedBgEnd+"\n" {
		t.Error("RedBgln failed")
	}
}

func TestGreenBgln(t *testing.T) {
	output := captureOutput(func() {
		GreenBgln("test", "word")
	})

	if output != GreenBgStart+"test word"+GreenBgEnd+"\n" {
		t.Error("GreenBgln failed")
	}
}

func TestYellowBgln(t *testing.T) {
	output := captureOutput(func() {
		YellowBgln("test", "word")
	})

	if output != YellowBgStart+"test word"+YellowBgEnd+"\n" {
		t.Error("YellowBgln failed")
	}
}

func TestBlueBgln(t *testing.T) {
	output := captureOutput(func() {
		BlueBgln("test", "word")
	})

	if output != BlueBgStart+"test word"+BlueBgEnd+"\n" {
		t.Error("BlueBgln failed")
	}
}

func TestMagentaBgln(t *testing.T) {
	output := captureOutput(func() {
		MagentaBgln("test", "word")
	})

	if output != MagentaBgStart+"test word"+MagentaBgEnd+"\n" {
		t.Error("MagentaBgln failed")
	}
}

func TestTealBgln(t *testing.T) {
	output := captureOutput(func() {
		TealBgln("test", "word")
	})

	if output != TealBgStart+"test word"+TealBgEnd+"\n" {
		t.Error("TealBgln failed")
	}
}

func TestWhiteBgln(t *testing.T) {
	output := captureOutput(func() {
		WhiteBgln("test", "word")
	})

	if output != WhiteBgStart+"test word"+WhiteBgEnd+"\n" {
		t.Error("WhiteBgln failed")
	}
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}
