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
