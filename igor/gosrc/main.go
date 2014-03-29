package main

import "github.com/nsf/termbox-go"
import "fmt"
import "math/rand"
import "strconv"
import "time"

var WINDOW_WIDTH = 0
var WINDOW_HEIGHT = 0
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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
    WINDOW_WIDTH, WINDOW_HEIGHT = termbox.Size()
    termbox.Clear(termbox.ColorDefault, termbox.ColorRed)
    printf_tb((WINDOW_WIDTH / 2) - 8, 0, termbox.ColorCyan, termbox.ColorBlack, "--- I.G.O.R. ---")
    termbox.SetCell(0, 0, 0x253C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(WINDOW_WIDTH - 1, 0, 0x253C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(WINDOW_WIDTH - 1, WINDOW_HEIGHT - 1, 0x253C, termbox.ColorWhite, termbox.ColorBlack)
    termbox.SetCell(0, WINDOW_HEIGHT - 1, 0x253C, termbox.ColorWhite, termbox.ColorBlack)
    printf_tb(0, 1, termbox.ColorCyan, termbox.ColorBlack, strconv.Itoa(WINDOW_WIDTH))
    printf_tb(0, 2, termbox.ColorCyan, termbox.ColorBlack, strconv.Itoa(WINDOW_HEIGHT))
    rwidth := r.Intn(WINDOW_WIDTH)
    rwidth_c1 := r.Intn(rwidth)
    rwidth_c2 := r.Intn(WINDOW_WIDTH - rwidth) + rwidth
    rheight := r.Intn(WINDOW_HEIGHT)
    rheight_c1 := r.Intn(rheight)
    rheight_c2 := r.Intn(WINDOW_HEIGHT - rheight) + rheight
    printf_tb(0, 4, termbox.ColorCyan, termbox.ColorBlack, strconv.Itoa(rwidth))
    printf_tb(0, 3, termbox.ColorCyan, termbox.ColorBlack, strconv.Itoa(rheight))
    for j := 0; j < WINDOW_WIDTH; j++ {
      printf_tb(j,rheight,termbox.ColorCyan, termbox.ColorBlack, "-")
      printf_tb(j,rheight_c1,termbox.ColorCyan, termbox.ColorBlack, "-")
      printf_tb(j,rheight_c2,termbox.ColorCyan, termbox.ColorBlack, "-")
    }
    for i := 0; i < WINDOW_HEIGHT; i++ {
      printf_tb(rwidth,i,termbox.ColorCyan, termbox.ColorBlack, "|")
      printf_tb(rwidth_c1,i,termbox.ColorCyan, termbox.ColorBlack, "|")
      printf_tb(rwidth_c2,i,termbox.ColorCyan, termbox.ColorBlack, "|")
    }
    printf_tb(rwidth,rheight,termbox.ColorRed, termbox.ColorBlack, "*")
    printf_tb(rwidth,rheight-2,termbox.ColorRed, termbox.ColorBlack, "N")
    printf_tb(rwidth,rheight+2,termbox.ColorRed, termbox.ColorBlack, "S")
    printf_tb(rwidth+3,rheight,termbox.ColorRed, termbox.ColorBlack, "E")
    printf_tb(rwidth-3,rheight,termbox.ColorRed, termbox.ColorBlack, "W")
    printf_tb(rwidth-2,rheight-1,termbox.ColorRed, termbox.ColorBlack, "NW")
    printf_tb(rwidth+1,rheight-1,termbox.ColorRed, termbox.ColorBlack, "NE")
    printf_tb(rwidth-2,rheight+1,termbox.ColorRed, termbox.ColorBlack, "SW")
    printf_tb(rwidth+1,rheight+1,termbox.ColorRed, termbox.ColorBlack, "SE")

    termbox.Flush()
    ev := termbox.PollEvent()
    if(ev.Key == termbox.KeyCtrlC) {
      break loop
    }
  }
}
