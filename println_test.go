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
