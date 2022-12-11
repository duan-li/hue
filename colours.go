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

	BlackBg   = "\033[1;40m%s\033[0m"
	RedBg     = "\033[1;41m%s\033[0m" // error
	GreenBg   = "\033[1;42m%s\033[0m" // success
	YellowBg  = "\033[1;43m%s\033[0m" // warn
	BlueBg    = "\033[1;44m%s\033[0m" // info
	MagentaBg = "\033[1;45m%s\033[0m" // fatal
	TealBg    = "\033[1;46m%s\033[0m"
	WhiteBg   = "\033[1;47m%s\033[0m"
)
