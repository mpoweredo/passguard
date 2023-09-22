package utils

import (
	"fmt"
	"github.com/pkg/term"
	"log"
)

const (
	UP         byte = 65
	DOWN       byte = 66
	ESCAPE     byte = 27
	ENTER      byte = 13
	START      byte = 115
	BACKSPACE  byte = 127
	LEFT_ARROW byte = 0
)

var keys = map[byte]bool{
	UP:   true,
	DOWN: true,
}

func GetInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	t.Restore()
	t.Close()

	// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
	// The third byte is the key specific value we are looking for.
	// For example the left arrow key is '<esc>[A' while the right is '<esc>[C'
	// See: https://en.wikipedia.org/wiki/ANSI_escape_code
	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}

	return 0
}

func ClearCmd() {
	fmt.Print("\033[H\033[2J")
}
