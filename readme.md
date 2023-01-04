# hue
A package to directly print colored text with or without colored background to the terminal.


## Installation

```bash
go get github.com/duan-li/hub
# if you have go modules enabled, you can update your go.mod file with
go get -u github.com/duan-li/hub
```

## Usage

### Generate colorful string

```golang

str := hue.Sblack("this", "is", "black", "text") // thisisblacktext
str := hue.Sred("this", "is", "red", "text") // thisisredtext
str := hue.Sgreen("this", "is", "green", "text") // thisisgreentext
str := hue.Syellow("this", "is", "yellow", "text") // thisisyellowtext
str := hue.Sblue("this", "is", "blue", "text") // thisisbluetext
str := hue.Smagenta("this", "is", "magenta", "text") // thisismagentatext
str := hue.Steal("this", "is", "teal", "text") // thisistealtext
str := hue.Swhite("this", "is", "white", "text") // thisiswhitetext
str := hue.SblackBg("this", "is", "black", "background") // thisisblackbackground
str := hue.SredBg("this", "is", "red", "background") // thisisredbackground
str := hue.SgreenBg("this", "is", "green", "background") // thisisgreenbackground
str := hue.SyellowBg("this", "is", "yellow", "background") // thisisyellowbackground
str := hue.SblueBg("this", "is", "blue", "background") // thisisbluebackground
str := hue.SmagentaBg("this", "is", "magenta", "background") // thisismagentabackground
str := hue.StealBg("this", "is", "teal", "background") // thisistealbackground
str := hue.SwhiteBg("this", "is", "white", "background") // thisiswhitebackground

```



### Print format colored text

```go
str := hue.Sblackf("this is %s %s", "black", "text") // this is black text
str := hue.Sredf("this is %s %s", "red", "text") // this is red text
str := hue.Sgreenf("this is %s %s", "green", "text") // this is green text
str := hue.Syellowf("this is %s %s", "yellow", "text") // this is yellow text
str := hue.Sbluef("this is %s %s", "blue", "text") // this is blue text
str := hue.Smagentaf("this is %s %s", "magenta", "text") // this is magenta text
str := hue.Stealf("this is %s %s", "teal", "text") // this is teal text
str := hue.Swhitef("this is %s %s", "white", "text") // this is white text
str := hue.SblackBgf("this is %s %s", "black", "background") // this is black background
str := hue.SredBgf("this is %s %s", "red", "background") // this is red background
str := hue.SgreenBgf("this is %s %s", "green", "background") // this is green background
str := hue.SyellowBgf("this is %s %s", "yellow", "background") // this is yellow background
str := hue.SblueBgf("this is %s %s", "blue", "background") // this is blue background
str := hue.SmagentaBgf("this is %s %s", "magenta", "background") // this is magenta background
str := hue.StealBgf("this is %s %s", "teal", "background") // this is teal background
str := hue.SwhiteBgf("this is %s %s", "white", "background") // this is white background
```

### Print a line of colored text

```go
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
```

### Command line output

```go
hue.Info("this", "is", "info")
hue.Warn("this", "is", "warn")
hue.Error("this", "is", "error")
hue.Success("this", "is", "success")
hue.Fatal("this", "is", "fatal")
```