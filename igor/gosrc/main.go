package main

import "github.com/limetext/termbox-go"
import "fmt"
import "./igor"
import "math/rand"
//import "strconv"

var COLOR [256]termbox.Attribute
var MAP *igor.Map


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
  termbox.Clear(COLOR[0], COLOR[rand.Intn(len(COLOR))])
loop:
	for {
    igor.WinWidth, igor.WinHeight = termbox.Size()
    MAP = igor.Clear(igor.WinWidth, igor.WinHeight)
    sections := igor.Split(igor.WinWidth, igor.WinHeight)
    for i := 0; i < 99; i++ {
      s1 := sections[rand.Intn(len(sections))]
      s2 := sections[rand.Intn(len(sections))]
      igor.Connect(s1, s2, MAP)
    }
    for x := 0; x < len(MAP.Tiles); x++ {
      row := MAP.Tiles[x]
      for y := 0; y < len(MAP.Tiles[x]); y++ {
        t := row[y]
        printf_tb(t.X, t.Y, COLOR[rand.Intn(len(COLOR))], COLOR[0], t.I)
      }
    }
    printf_tb((igor.WinWidth/2)-8, 0, COLOR[32], COLOR[0], "--- I.G.O.R. ---")
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

/*
		termbox.SetCell(0, 0, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(igor.WinWidth-1, 0, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(igor.WinWidth-1, igor.WinHeight-1, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(0, igor.WinHeight-1, 0x253C, COLOR[255], COLOR[0])

		printf_tb(0, 1, COLOR[32], COLOR[0], strconv.Itoa(igor.WinWidth))
		printf_tb(0, 2, COLOR[32], COLOR[0], strconv.Itoa(igor.WinHeight))
		printf_tb(0, 4, COLOR[32], COLOR[0], strconv.Itoa(rwidth))
		printf_tb(0, 3, COLOR[32], COLOR[0], strconv.Itoa(rheight))

		currentWidth := 0
		currentHeight := rheight

		for i := 0; i < igor.WinWidth; i++ {
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
		for j := 0; j < igor.WinWidth; j++ {
			printf_tb(j, rheight, COLOR[16], COLOR[0], "-")
			printf_tb(j, rheight_c1, COLOR[16], COLOR[0], "-")
			printf_tb(j, rheight_c2, COLOR[16], COLOR[0], "-")
		}

		for i := 0; i < igor.WinHeight; i++ {
			printf_tb(rwidth, i, COLOR[16], COLOR[0], "|")
			printf_tb(rwidth_c1, i, COLOR[16], COLOR[0], "|")
			printf_tb(rwidth_c2, i, COLOR[16], COLOR[0], "|")
		}

		printf_tb(0, rheight_c1-1, COLOR[255], COLOR[0], "01")
		printf_tb(rwidth_c1+1, rheight_c1-1, COLOR[255], COLOR[0], "02")
		printf_tb(rwidth+1, rheight_c1-1, COLOR[255], COLOR[0], "03")
		printf_tb(rwidth_c2+1, rheight_c1-1, COLOR[255], COLOR[0], "04")
		printf_tb(igor.WinWidth-2, rheight_c1-1, COLOR[255], COLOR[0], "05")

		printf_tb(0, rheight-1, COLOR[255], COLOR[0], "06")
		printf_tb(rwidth_c1+1, rheight-1, COLOR[255], COLOR[0], "07")
		printf_tb(rwidth+1, rheight-1, COLOR[255], COLOR[0], "08")
		printf_tb(rwidth_c2+1, rheight-1, COLOR[255], COLOR[0], "09")
		printf_tb(igor.WinWidth-2, rheight-1, COLOR[255], COLOR[0], "10")

		printf_tb(0, rheight_c2-1, COLOR[255], COLOR[0], "11")
		printf_tb(rwidth_c1+1, rheight_c2-1, COLOR[255], COLOR[0], "12")
		printf_tb(rwidth+1, rheight_c2-1, COLOR[255], COLOR[0], "13")
		printf_tb(rwidth_c2+1, rheight_c2-1, COLOR[255], COLOR[0], "14")
		printf_tb(igor.WinWidth-2, rheight_c2-1, COLOR[255], COLOR[0], "15")

		printf_tb(0, igor.WinHeight-1, COLOR[255], COLOR[0], "16")
		printf_tb(rwidth_c1+1, igor.WinHeight-1, COLOR[255], COLOR[0], "17")
		printf_tb(rwidth+1, igor.WinHeight-1, COLOR[255], COLOR[0], "18")
		printf_tb(rwidth_c2+1, igor.WinHeight-1, COLOR[255], COLOR[0], "19")
		printf_tb(igor.WinWidth-2, igor.WinHeight-1, COLOR[255], COLOR[0], "20")

		printf_tb(rwidth, rheight, COLOR[120], COLOR[0], "*")
		printf_tb(rwidth, rheight-2, COLOR[120], COLOR[0], "N")
		printf_tb(rwidth, rheight+2, COLOR[120], COLOR[0], "S")
		printf_tb(rwidth+3, rheight, COLOR[120], COLOR[0], "E")
		printf_tb(rwidth-3, rheight, COLOR[120], COLOR[0], "W")
		printf_tb(rwidth-2, rheight-1, COLOR[120], COLOR[0], "NW")
		printf_tb(rwidth+1, rheight-1, COLOR[120], COLOR[0], "NE")
		printf_tb(rwidth-2, rheight+1, COLOR[120], COLOR[0], "SW")
		printf_tb(rwidth+1, rheight+1, COLOR[120], COLOR[0], "SE")

*/
