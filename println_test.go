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
