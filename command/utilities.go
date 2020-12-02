package command

import (
	"fmt"

	"github.com/jedib0t/go-pretty/text"
)

func line(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

func lineColor(color text.Color, format string, a ...interface{}) {
	fmt.Println(text.Colors{color}.Sprintf(format, a...))
}
