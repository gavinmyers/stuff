package main

import "github.com/limetext/termbox-go"
import "fmt"
import "math/rand"
import "strconv"

var WINDOW_WIDTH = 0
var WINDOW_HEIGHT = 0
var COLOR [256]termbox.Attribute

func init() {
	err := termbox.Init()
	termbox.SetColorMode(termbox.ColorMode256)
	if err != nil {
		panic(err)
	}
  for i := 0; i < 255; i++ {
    COLOR[i] = termbox.Attribute(i)
  }
}

func main() {
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(COLOR[0], COLOR[255])
loop:
	for {
		WINDOW_WIDTH, WINDOW_HEIGHT = termbox.Size()
		termbox.Clear(COLOR[0], COLOR[rand.Intn(len(COLOR))])

		termbox.SetCell(0, 0, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(WINDOW_WIDTH-1, 0, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(WINDOW_WIDTH-1, WINDOW_HEIGHT-1, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(0, WINDOW_HEIGHT-1, 0x253C, COLOR[255], COLOR[0])


		rwidth := rand.Intn(WINDOW_WIDTH/2) + WINDOW_WIDTH/4

		rwidth_c1 := rand.Intn(rwidth/2) + rwidth/4
		rwidth_c2 := rand.Intn((WINDOW_WIDTH-rwidth)/2) +
                 rwidth + ((WINDOW_WIDTH - rwidth) / 4)

		rheight := rand.Intn(WINDOW_HEIGHT/2) + WINDOW_HEIGHT/4

		rheight_c1 := rand.Intn(rheight/2) + rheight/4
		rheight_c2 := rand.Intn((WINDOW_HEIGHT-rheight)/2) +
                  rheight + ((WINDOW_HEIGHT - rheight) / 4)

		printf_tb(0, 1, COLOR[32], COLOR[0], strconv.Itoa(WINDOW_WIDTH))
		printf_tb(0, 2, COLOR[32], COLOR[0], strconv.Itoa(WINDOW_HEIGHT))
		printf_tb(0, 4, COLOR[32], COLOR[0], strconv.Itoa(rwidth))
		printf_tb(0, 3, COLOR[32], COLOR[0], strconv.Itoa(rheight))

		printf_tb(0, rheight_c1-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth+1, rheight_c1-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c1+1, rheight_c1-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c2+1, rheight_c1-1, COLOR[255], COLOR[0], ".")
		printf_tb(WINDOW_WIDTH-1, rheight_c1-1, COLOR[255], COLOR[0], ".")

		printf_tb(0, rheight-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth+1, rheight-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c1+1, rheight-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c2+1, rheight-1, COLOR[255], COLOR[0], ".")
		printf_tb(WINDOW_WIDTH-1, rheight-1, COLOR[255], COLOR[0], ".")

		printf_tb(0, rheight_c2-1, COLOR[rand.Intn(len(COLOR))], COLOR[0], "*")

		currentWidth := 0
		currentHeight := rheight

		for i := 0; i < WINDOW_WIDTH; i++ {
      termbox.SetCell(currentWidth,
                      currentHeight,
                      0x00A4,
                      COLOR[rand.Intn(len(COLOR))],
                      COLOR[rand.Intn(len(COLOR))])
			if rand.Intn(2) == 1 {
				currentHeight--
			} else {
				currentHeight++
			}
			if rand.Intn(2) == 1 {
				currentWidth--
			} else {
				currentWidth++
			}
		}

		currentWidth = 0
		currentHeight = rheight
		for i := 0; i < WINDOW_WIDTH; i++ {
      termbox.SetCell(currentWidth,
                      currentHeight,
                      0x00A4,
                      COLOR[rand.Intn(len(COLOR))],
                      COLOR[rand.Intn(len(COLOR))])
			if rand.Intn(3) == 1 {
				currentHeight--
			} else {
				currentHeight++
			}
			if rand.Intn(3) == 1 {
				currentWidth--
			} else {
				currentWidth++
			}
		}

		printf_tb(rwidth+1, rheight_c2-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c1+1, rheight_c2-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c2+1, rheight_c2-1, COLOR[255], COLOR[0], ".")
		printf_tb(WINDOW_WIDTH-1, rheight_c2-1, COLOR[255], COLOR[0], ".")

		printf_tb(0, WINDOW_HEIGHT-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth+1, WINDOW_HEIGHT-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c1+1, WINDOW_HEIGHT-1, COLOR[255], COLOR[0], ".")
		printf_tb(rwidth_c2+1, WINDOW_HEIGHT-1, COLOR[255], COLOR[0], ".")
		printf_tb(WINDOW_WIDTH-1, WINDOW_HEIGHT-1, COLOR[255], COLOR[0], ".")

		for j := 0; j < WINDOW_WIDTH; j++ {
			printf_tb(j, rheight, COLOR[16], COLOR[0], "-")
			printf_tb(j, rheight_c1, COLOR[16], COLOR[0], "-")
			printf_tb(j, rheight_c2, COLOR[16], COLOR[0], "-")
		}

		for i := 0; i < WINDOW_HEIGHT; i++ {
			printf_tb(rwidth, i, COLOR[16], COLOR[0], "|")
			printf_tb(rwidth_c1, i, COLOR[16], COLOR[0], "|")
			printf_tb(rwidth_c2, i, COLOR[16], COLOR[0], "|")
		}

		printf_tb(rwidth, rheight, COLOR[120], COLOR[0], "*")
		printf_tb(rwidth, rheight-2, COLOR[120], COLOR[0], "N")
		printf_tb(rwidth, rheight+2, COLOR[120], COLOR[0], "S")
		printf_tb(rwidth+3, rheight, COLOR[120], COLOR[0], "E")
		printf_tb(rwidth-3, rheight, COLOR[120], COLOR[0], "W")
		printf_tb(rwidth-2, rheight-1, COLOR[120], COLOR[0], "NW")
		printf_tb(rwidth+1, rheight-1, COLOR[120], COLOR[0], "NE")
		printf_tb(rwidth-2, rheight+1, COLOR[120], COLOR[0], "SW")
		printf_tb(rwidth+1, rheight+1, COLOR[120], COLOR[0], "SE")

		printf_tb((WINDOW_WIDTH/2)-8, 0, COLOR[32], COLOR[0], "--- I.G.O.R. ---")
		termbox.Flush()
		ev := termbox.PollEvent()
		if ev.Key == termbox.KeyCtrlC {
			break loop
		}
	}
}

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


