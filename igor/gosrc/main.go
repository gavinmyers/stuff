package main

import "github.com/limetext/termbox-go"
import "./igor"
import "./gruyere"
import "./moo"
import "math/rand"

var WORLD moo.World
var MAP *igor.Map


func init() {
  gruyere.Init()
}

func main() {
	defer termbox.Close()
  gruyere.Clear()
  igor.WinWidth, igor.WinHeight = gruyere.Size()
  sections := igor.Split(igor.WinWidth, igor.WinHeight)
  MAP = igor.Clear(igor.WinWidth, igor.WinHeight)
  for i := 0; i < 99; i++ {
    s1 := sections[rand.Intn(len(sections))]
    s2 := sections[rand.Intn(len(sections))]
    igor.Connect(s1, s2, MAP)
  }

loop:
	for {
    for x := 0; x < len(MAP.Tiles); x++ {
      row := MAP.Tiles[x]
      for y := 0; y < len(MAP.Tiles[x]); y++ {
        t := row[y]
        gruyere.Draw(t.X, t.Y, gruyere.Color[rand.Intn(len(gruyere.Color))], gruyere.Color[0], t.I)
      }
    }
    gruyere.Draw((igor.WinWidth/2)-8, 0, gruyere.Color[32], gruyere.Color[0], "--- I.G.O.R. ---")
    gruyere.Flush()
    ev := termbox.PollEvent()
    if ev.Key == termbox.KeyCtrlC {
      break loop
    }
  }
}


