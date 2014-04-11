package main

import "./moo"
import "github.com/limetext/termbox-go"

func main() {
	defer termbox.Close()
  sb := &moo.SpriteBuilder {}
  sb.Init()
  gui := &moo.GUI {}
  gui.Init()
loop:
	for {
    gui.Draw(25,25,gui.Color[0],gui.Color[255],"@")

    termbox.Flush()
    ev := termbox.PollEvent()
    if ev.Key == termbox.KeyCtrlC {
      break loop
    }
  }

}
