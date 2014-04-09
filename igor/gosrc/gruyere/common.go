package gruyere

import "github.com/limetext/termbox-go"
import "fmt"


func Init() {
  err := termbox.Init()
	termbox.SetColorMode(termbox.ColorMode256)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 255; i++ {
		Color[i] = termbox.Attribute(i)
	}
}

func Clear() {
  termbox.Clear(Color[0], Color[0])
}


func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func Draw(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}


