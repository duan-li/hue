package hue

// Constant of colours and background colours
const (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m" // error
	Green   = "\033[1;32m%s\033[0m" // success
	Yellow  = "\033[1;33m%s\033[0m" // warn
	Blue    = "\033[1;34m%s\033[0m" // info
	Magenta = "\033[1;35m%s\033[0m" // fatal
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"

	BlackStart   = "\033[1;30m"
	RedStart     = "\033[1;31m" // error
	GreenStart   = "\033[1;32m" // success
	YellowStart  = "\033[1;33m" // warn
	BlueStart    = "\033[1;34m" // info
	MagentaStart = "\033[1;35m" // fatal
	TealStart    = "\033[1;36m"
	WhiteStart   = "\033[1;37m"

	BlackEnd   = "\033[0m"
	RedEnd     = "\033[0m" // error
	GreenEnd   = "\033[0m" // success
	YellowEnd  = "\033[0m" // warn
	BlueEnd    = "\033[0m" // info
	MagentaEnd = "\033[0m" // fatal
	TealEnd    = "\033[0m"
	WhiteEnd   = "\033[0m"

	BlackBg   = "\033[1;40m%s\033[0m"
	RedBg     = "\033[1;41m%s\033[0m" // error
	GreenBg   = "\033[1;42m%s\033[0m" // success
	YellowBg  = "\033[1;43m%s\033[0m" // warn
	BlueBg    = "\033[1;44m%s\033[0m" // info
	MagentaBg = "\033[1;45m%s\033[0m" // fatal
	TealBg    = "\033[1;46m%s\033[0m"
	WhiteBg   = "\033[1;47m%s\033[0m"

	BlackBgStart   = "\033[1;40m"
	RedBgStart     = "\033[1;41m" // error
	GreenBgStart   = "\033[1;42m" // success
	YellowBgStart  = "\033[1;43m" // warn
	BlueBgStart    = "\033[1;44m" // info
	MagentaBgStart = "\033[1;45m" // fatal
	TealBgStart    = "\033[1;46m"
	WhiteBgStart   = "\033[1;47m"

	BlackBgEnd   = "\033[0m"
	RedBgEnd     = "\033[0m" // error
	GreenBgEnd   = "\033[0m" // success
	YellowBgEnd  = "\033[0m" // warn
	BlueBgEnd    = "\033[0m" // info
	MagentaBgEnd = "\033[0m" // fatal
	TealBgEnd    = "\033[0m"
	WhiteBgEnd   = "\033[0m"
)
