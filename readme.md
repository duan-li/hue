# hue
A package to directly print colored text with or without colored background to the terminal.

There are many ways to make a console command output more colorful. One way to do this on a Linux or Mac terminal is to use ANSI escape codes. These codes allow you to change the color of text, background, and cursor in the terminal.

## Installation

To install a package in Go, use the "go get" command followed by the package name. For example:

```bash
go get github.com/duan-li/hub
# if you have go modules enabled, you can update your go.mod file with
go get -u github.com/duan-li/hub
```

This will download the package and install it in your Go workspace. You can then import the package into your code using an import statement like this:

```go
import "github.com/duan-li/hub"
```

If you want to update an installed package to the latest version, use the "go get -u" command. For example:

```go
go get -u github.com/duan-li/hub
```


## Usage

To use a package in a Go program, you first need to import it. Importing a package allows you to access the types, functions, and variables that it defines. You can import a package using the import keyword, followed by the package name in double quotes. For example, to import the fmt package, which provides basic input and output functions, you would write:

```go
import "github.com/duan-li/hub"
```

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
hue.Blackln("this", "is", "black", "line") // this is black line
hue.Redln("this", "is", "red", "line") // this is red line
hue.Greenln("this", "is", "green", "line") // this is green line
hue.Yellowln("this", "is", "yellow", "line") // this is yellow line
hue.Blueln("this", "is", "blue", "line") // this is blue line
hue.Magentaln("this", "is", "magenta", "line") // this is magenta line
hue.Tealln("this", "is", "teal", "line") // this is teal line
hue.Whiteln("this", "is", "white", "line") // this is white line
hue.BlackBgln("this", "is", "black", "background", "line") // this is black background line
hue.RedBgln("this", "is", "red", "background", "line") // this is red background line
hue.GreenBgln("this", "is", "green", "background", "line") // this is green background line
hue.YellowBgln("this", "is", "yellow", "background", "line") // this is yellow background line
hue.BlueBgln("this", "is", "blue", "background", "line") // this is blue background line
hue.MagentaBgln("this", "is", "magenta", "background", "line") // this is magenta background line
hue.TealBgln("this", "is", "teal", "background", "line") // this is teal background line
hue.WhiteBgln("this", "is", "white", "background", "line") // this is white background line
```

### Command line output

```go
hue.Info("this", "is", "info") // this is info 
hue.Warn("this", "is", "warn") // this is warn 
hue.Error("this", "is", "error") // this is error 
hue.Success("this", "is", "success") // this is success 
hue.Fatal("this", "is", "fatal") // this is fatal 
```