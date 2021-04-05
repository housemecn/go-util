package log

import (
	"fmt"
	"log"
)

// Color list
const (
	Red = uint8(iota + 31)
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// LW log with color and wrap
func LW(color uint8, s string) {
	L(color, s+"\n")
}

// L log with color
func L(color uint8, s string) {
	tmp := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
	if len(tmp) > 0 {
		log.Print(tmp)
	}
}
