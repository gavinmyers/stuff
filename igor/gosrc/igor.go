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
    gui.Draw(25,25,gui.Color[125],gui.Color[25],"@")
    gui.Flush()
    key := gui.PollEvent()
    if key == termbox.KeyCtrlC {
      break loop
    }
  }

}
