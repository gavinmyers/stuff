package main

import "github.com/limetext/termbox-go"
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
  termbox.SetColorMode(termbox.ColorMode256)
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  termbox.SetInputMode(termbox.InputEsc)
  termbox.Clear(termbox.Attribute(0), termbox.Attribute(255))
loop:
  for {
    WINDOW_WIDTH, WINDOW_HEIGHT = termbox.Size()
    termbox.Clear(termbox.ColorDefault, termbox.Attribute(120))
    printf_tb((WINDOW_WIDTH / 2) - 8, 0, termbox.Attribute(32), termbox.Attribute(0), "--- I.G.O.R. ---")
    termbox.SetCell(0, 0, 0x253C, termbox.Attribute(255), termbox.Attribute(0))
    termbox.SetCell(WINDOW_WIDTH - 1, 0, 0x253C, termbox.Attribute(255), termbox.Attribute(0))
    termbox.SetCell(WINDOW_WIDTH - 1, WINDOW_HEIGHT - 1, 0x253C, termbox.Attribute(255), termbox.Attribute(0))
    termbox.SetCell(0, WINDOW_HEIGHT - 1, 0x253C, termbox.Attribute(255), termbox.Attribute(0))
    printf_tb(0, 1, termbox.Attribute(32), termbox.Attribute(0), strconv.Itoa(WINDOW_WIDTH))
    printf_tb(0, 2, termbox.Attribute(32), termbox.Attribute(0), strconv.Itoa(WINDOW_HEIGHT))
    rwidth := r.Intn(WINDOW_WIDTH / 2) + WINDOW_WIDTH / 4
    rwidth_c1 := r.Intn(rwidth / 2) + rwidth / 4
    rwidth_c2 := r.Intn((WINDOW_WIDTH - rwidth) / 2) + rwidth + ((WINDOW_WIDTH - rwidth) / 4)
    rheight := r.Intn(WINDOW_HEIGHT / 2) + WINDOW_HEIGHT / 4
    rheight_c1 := r.Intn(rheight / 2) + rheight / 4
    rheight_c2 := r.Intn((WINDOW_HEIGHT - rheight) / 2) + rheight + ((WINDOW_HEIGHT - rheight) / 4)
    printf_tb(0, 4, termbox.Attribute(32), termbox.Attribute(0), strconv.Itoa(rwidth))
    printf_tb(0, 3, termbox.Attribute(32), termbox.Attribute(0), strconv.Itoa(rheight))

    printf_tb(0,rheight_c1-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth+1,rheight_c1-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c1+1,rheight_c1-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c2+1,rheight_c1-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(WINDOW_WIDTH-1,rheight_c1-1,termbox.Attribute(255), termbox.Attribute(0), ".")

    printf_tb(0,rheight-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth+1,rheight-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c1+1,rheight-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c2+1,rheight-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(WINDOW_WIDTH-1,rheight-1,termbox.Attribute(255), termbox.Attribute(0), ".")

    printf_tb(0,rheight_c2-1,termbox.ColorMagenta, termbox.Attribute(0), "*")
    currentWidth := 0
    currentHeight := rheight
    for i := 0; i < WINDOW_WIDTH; i++ {
      printf_tb(currentWidth,currentHeight,termbox.ColorMagenta, termbox.Attribute(0), "*")
      if(r.Intn(3) == 1) {
        currentHeight--
      } else if(r.Intn(3) == 1) {
        currentHeight++
      } else {
        currentHeight++
      }
      if(r.Intn(3) == 1) {
        currentWidth++
      } else if(r.Intn(3) == 1) {
        currentWidth++
      } else {
        currentWidth++
      }
    }
    currentWidth = 0
    currentHeight = rheight
    for i := 0; i < WINDOW_WIDTH; i++ {
      printf_tb(currentWidth,currentHeight,termbox.ColorMagenta, termbox.Attribute(0), "*")
      if(r.Intn(3) == 1) {
        currentHeight--
      } else if(r.Intn(3) == 1) {
        currentHeight++
      } else {
        currentHeight++
      }
      if(r.Intn(3) == 1) {
        currentWidth++
      } else if(r.Intn(3) == 1) {
        currentWidth++
      } else {
        currentWidth++
      }
    }


    printf_tb(rwidth+1,rheight_c2-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c1+1,rheight_c2-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c2+1,rheight_c2-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(WINDOW_WIDTH-1,rheight_c2-1,termbox.Attribute(255), termbox.Attribute(0), ".")

    printf_tb(0,WINDOW_HEIGHT-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth+1,WINDOW_HEIGHT-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c1+1,WINDOW_HEIGHT-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(rwidth_c2+1,WINDOW_HEIGHT-1,termbox.Attribute(255), termbox.Attribute(0), ".")
    printf_tb(WINDOW_WIDTH-1,WINDOW_HEIGHT-1,termbox.Attribute(255), termbox.Attribute(0), ".")

    for j := 0; j < WINDOW_WIDTH; j++ {
      printf_tb(j,rheight,termbox.Attribute(16), termbox.Attribute(0), "-")
      printf_tb(j,rheight_c1,termbox.Attribute(16), termbox.Attribute(0), "-")
      printf_tb(j,rheight_c2,termbox.Attribute(16), termbox.Attribute(0), "-")
    }
    for i := 0; i < WINDOW_HEIGHT; i++ {
      printf_tb(rwidth,i,termbox.Attribute(16), termbox.Attribute(0), "|")
      printf_tb(rwidth_c1,i,termbox.Attribute(16), termbox.Attribute(0), "|")
      printf_tb(rwidth_c2,i,termbox.Attribute(16), termbox.Attribute(0), "|")
    }
    printf_tb(rwidth,rheight,termbox.Attribute(120), termbox.Attribute(0), "*")
    printf_tb(rwidth,rheight-2,termbox.Attribute(120), termbox.Attribute(0), "N")
    printf_tb(rwidth,rheight+2,termbox.Attribute(120), termbox.Attribute(0), "S")
    printf_tb(rwidth+3,rheight,termbox.Attribute(120), termbox.Attribute(0), "E")
    printf_tb(rwidth-3,rheight,termbox.Attribute(120), termbox.Attribute(0), "W")
    printf_tb(rwidth-2,rheight-1,termbox.Attribute(120), termbox.Attribute(0), "NW")
    printf_tb(rwidth+1,rheight-1,termbox.Attribute(120), termbox.Attribute(0), "NE")
    printf_tb(rwidth-2,rheight+1,termbox.Attribute(120), termbox.Attribute(0), "SW")
    printf_tb(rwidth+1,rheight+1,termbox.Attribute(120), termbox.Attribute(0), "SE")

    termbox.Flush()
    ev := termbox.PollEvent()
    if(ev.Key == termbox.KeyCtrlC) {
      break loop
    }
  }
}
