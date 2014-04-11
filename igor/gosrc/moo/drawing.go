package moo

import "github.com/limetext/termbox-go"
import "fmt"

type GUI struct {
  Width int
  Height int
  Color [256]termbox.Attribute
}

func (c *GUI) Init() {
  err := termbox.Init()
	termbox.SetColorMode(termbox.ColorMode256)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 255; i++ {
		c.Color[i] = termbox.Attribute(i)
	}
	termbox.SetInputMode(termbox.InputEsc)
  termbox.Clear(c.Color[0], c.Color[0])
  c.Width, c.Height = termbox.Size()
}

func (c *GUI) Draw(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	c.print_tb(x, y, fg, bg, s)
}

func (c *GUI) print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}


