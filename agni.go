package agni

import (
	"fmt"
	"log"
	"strings"
)

const (
	BlackText   = "\033[30m"
	RedText     = "\033[31m"
	GreenText   = "\033[32m"
	YellowText  = "\033[33m"
	BlueText    = "\033[34m"
	MagentaText = "\033[35m"
	CyanText    = "\033[36m"
	WhiteText   = "\033[37m"

	BlackBackground   = "\033[40m"
	RedBackground     = "\033[41m"
	GreenBackground   = "\033[42m"
	YellowBackground  = "\033[43m"
	BlueBackground    = "\033[44m"
	MagentaBackground = "\033[45m"
	CyanBackground    = "\033[46m"
	WhiteBackground   = "\033[47m"

	Bold       = "\033[1m"
	Underline  = "\033[4m"
	Reverse    = "\033[7m"
	ResetStyle = "\033[0m"
)

func Logger(msg interface{}, style ...string) {
	msg = fmt.Sprintf("%s%s"+ResetStyle, resolveStyles(style), msg)

	log.Println(msg)
}

func Println(msg interface{}, style ...string) {
	msg = fmt.Sprintf("%s%s"+ResetStyle, resolveStyles(style), msg)

	fmt.Println(msg)
}

func resolveStyles(styles []string) string {
	return strings.Join(styles, "")
}
