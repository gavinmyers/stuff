package main

import "github.com/nsf/termbox-go"
import "fmt"

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
  for _, c := range msg {
    termbox.SetCell(x, y, c, fg, bg)
    x++
  }
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
  s := fmt.Sprintf(format, args...)
  print_tb(x, y, fg, bg, s)
}


func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
loop:
  for {
    printf_tb(2, 20, termbox.ColorCyan, termbox.ColorBlack, "@")
    printf_tb(4, 20, termbox.ColorCyan, termbox.ColorBlack, "Hello World")
    termbox.SetCell(0, 0, 0x250C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(1, 1, 0x251C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(2, 2, 0x252C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(2, 3, 0x165C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(4, 5, 0x0298, termbox.ColorWhite, termbox.ColorBlack)
    termbox.Flush()
    ev := termbox.PollEvent()
    if(ev.Key == termbox.KeyCtrlC) {
      break loop
    }
  }
}
