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
    printf_tb(3, 19, termbox.ColorWhite, termbox.ColorBlack, "Key: ")
    termbox.Flush()
    ev := termbox.PollEvent()
    if(ev.Key == termbox.KeyCtrlC) {
      break loop
    }
  }
}
