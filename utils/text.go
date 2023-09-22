package utils

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

func CenterText(s string) *bytes.Buffer {
	fd := int(os.Stdin.Fd())
	w, _, err := term.GetSize(fd)

	if err != nil {
		w = 0
	}

	const half, space = 2, "\u0020"
	var b bytes.Buffer
	n := (w - utf8.RuneCountInString(s)) / half
	if n < 1 {
		fmt.Fprint(&b, s)
	}
	fmt.Fprintf(&b, "%s%s", strings.Repeat(space, n), s)
	return &b
}

func CreateHorizontalLine() {
	color.Set(color.FgHiGreen)
	defer color.Unset()

	char := "-"

	fd := int(os.Stdin.Fd())
	w, _, err := term.GetSize(fd)
	s := char

	if err != nil {
		w = 1
	}

	s = strings.Repeat(char, w)

	fmt.Println(s)
}

func PrintWithColor(m any, o ...color.Attribute) {
	c := color.FgWhite

	if len(o) == 1 {
		c = o[0]
	}

	color.Set(c)
	defer color.Unset()

	fmt.Println(m)
}
