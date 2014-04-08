package main

import "github.com/limetext/termbox-go"
import "fmt"
import "./igor"
import "math/rand"
import "strconv"

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
loop:
	for {

    igor.WINDOW_WIDTH, igor.WINDOW_HEIGHT = termbox.Size()
    MAP = &igor.Map {Tiles:make([][]*igor.Tile, igor.WINDOW_WIDTH * 2)}
    for i := 0; i < igor.WINDOW_WIDTH * 2; i++ {
      MAP.Tiles[i] = make([]*igor.Tile, igor.WINDOW_HEIGHT * 2)
      for j := 0; j < igor.WINDOW_HEIGHT * 2; j++ {
        t := &igor.Tile {X:i, Y:j, I: "▒"}
        MAP.Tiles[i][j] = t
      }
    }

		termbox.Clear(COLOR[0], COLOR[rand.Intn(len(COLOR))])

		termbox.SetCell(0, 0, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(igor.WINDOW_WIDTH-1, 0, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(igor.WINDOW_WIDTH-1, igor.WINDOW_HEIGHT-1, 0x253C, COLOR[255], COLOR[0])
		termbox.SetCell(0, igor.WINDOW_HEIGHT-1, 0x253C, COLOR[255], COLOR[0])

		rwidth := rand.Intn(igor.WINDOW_WIDTH/2) + igor.WINDOW_WIDTH/4

		rwidth_c1 := rand.Intn(rwidth/2) + rwidth/4
		rwidth_c2 := rand.Intn((igor.WINDOW_WIDTH-rwidth)/2) +
			rwidth + ((igor.WINDOW_WIDTH - rwidth) / 4)

		rheight := rand.Intn(igor.WINDOW_HEIGHT/2) + igor.WINDOW_HEIGHT/4

		rheight_c1 := rand.Intn(rheight/2) + rheight/4
		rheight_c2 := rand.Intn((igor.WINDOW_HEIGHT-rheight)/2) +
			rheight + ((igor.WINDOW_HEIGHT - rheight) / 4)

    sections := make([][2]int, 25)
    sections[0][0] = 0
    sections[0][1] = rheight_c1
    sections[1][0] = rwidth_c1
    sections[1][1] = rheight_c1
    sections[2][0] = rwidth
    sections[2][1] = rheight_c1
    sections[3][0] = rwidth_c2
    sections[3][1] = rheight_c1
    sections[4][0] = igor.WINDOW_WIDTH
    sections[4][1] = rheight_c1
    sections[5][0] = 0
    sections[5][1] = rheight
    sections[6][0] = rwidth_c1
    sections[6][1] = rheight
    sections[7][0] = rwidth
    sections[7][1] = rheight
    sections[8][0] = rwidth_c2
    sections[8][1] = rheight
    sections[9][0] = igor.WINDOW_WIDTH
    sections[9][1] = rheight
    sections[10][0] = 0
    sections[10][1] = rheight_c2
    sections[11][0] = rwidth_c1
    sections[11][1] = rheight_c2
    sections[12][0] = rwidth
    sections[12][1] = rheight_c2
    sections[13][0] = rwidth_c2
    sections[13][1] = rheight_c2
    sections[14][0] = igor.WINDOW_WIDTH
    sections[14][1] = rheight_c2
    sections[15][0] = 0
    sections[15][1] = igor.WINDOW_HEIGHT
    sections[16][0] = rwidth_c1
    sections[16][1] = igor.WINDOW_HEIGHT
    sections[17][0] = rwidth
    sections[17][1] = igor.WINDOW_HEIGHT
    sections[18][0] = rwidth_c2
    sections[18][1] = igor.WINDOW_HEIGHT
    sections[19][0] = igor.WINDOW_WIDTH
    sections[19][1] = igor.WINDOW_HEIGHT
    sections[20][0] = 0
    sections[20][1] = 0
    sections[21][0] = rwidth_c1
    sections[21][1] = 0
    sections[22][0] = rwidth
    sections[22][1] = 0
    sections[23][0] = rwidth_c2
    sections[23][1] = 0
    sections[24][0] = igor.WINDOW_WIDTH
    sections[24][1] = 0

		printf_tb(0, 1, COLOR[32], COLOR[0], strconv.Itoa(igor.WINDOW_WIDTH))
		printf_tb(0, 2, COLOR[32], COLOR[0], strconv.Itoa(igor.WINDOW_HEIGHT))
		printf_tb(0, 4, COLOR[32], COLOR[0], strconv.Itoa(rwidth))
		printf_tb(0, 3, COLOR[32], COLOR[0], strconv.Itoa(rheight))

		currentWidth := 0
		currentHeight := rheight

		for i := 0; i < igor.WINDOW_WIDTH; i++ {
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
		for j := 0; j < igor.WINDOW_WIDTH; j++ {
			printf_tb(j, rheight, COLOR[16], COLOR[0], "-")
			printf_tb(j, rheight_c1, COLOR[16], COLOR[0], "-")
			printf_tb(j, rheight_c2, COLOR[16], COLOR[0], "-")
		}

		for i := 0; i < igor.WINDOW_HEIGHT; i++ {
			printf_tb(rwidth, i, COLOR[16], COLOR[0], "|")
			printf_tb(rwidth_c1, i, COLOR[16], COLOR[0], "|")
			printf_tb(rwidth_c2, i, COLOR[16], COLOR[0], "|")
		}

		printf_tb(0, rheight_c1-1, COLOR[255], COLOR[0], "01")
		printf_tb(rwidth_c1+1, rheight_c1-1, COLOR[255], COLOR[0], "02")
		printf_tb(rwidth+1, rheight_c1-1, COLOR[255], COLOR[0], "03")
		printf_tb(rwidth_c2+1, rheight_c1-1, COLOR[255], COLOR[0], "04")
		printf_tb(igor.WINDOW_WIDTH-2, rheight_c1-1, COLOR[255], COLOR[0], "05")

		printf_tb(0, rheight-1, COLOR[255], COLOR[0], "06")
		printf_tb(rwidth_c1+1, rheight-1, COLOR[255], COLOR[0], "07")
		printf_tb(rwidth+1, rheight-1, COLOR[255], COLOR[0], "08")
		printf_tb(rwidth_c2+1, rheight-1, COLOR[255], COLOR[0], "09")
		printf_tb(igor.WINDOW_WIDTH-2, rheight-1, COLOR[255], COLOR[0], "10")

		printf_tb(0, rheight_c2-1, COLOR[255], COLOR[0], "11")
		printf_tb(rwidth_c1+1, rheight_c2-1, COLOR[255], COLOR[0], "12")
		printf_tb(rwidth+1, rheight_c2-1, COLOR[255], COLOR[0], "13")
		printf_tb(rwidth_c2+1, rheight_c2-1, COLOR[255], COLOR[0], "14")
		printf_tb(igor.WINDOW_WIDTH-2, rheight_c2-1, COLOR[255], COLOR[0], "15")

		printf_tb(0, igor.WINDOW_HEIGHT-1, COLOR[255], COLOR[0], "16")
		printf_tb(rwidth_c1+1, igor.WINDOW_HEIGHT-1, COLOR[255], COLOR[0], "17")
		printf_tb(rwidth+1, igor.WINDOW_HEIGHT-1, COLOR[255], COLOR[0], "18")
		printf_tb(rwidth_c2+1, igor.WINDOW_HEIGHT-1, COLOR[255], COLOR[0], "19")
		printf_tb(igor.WINDOW_WIDTH-2, igor.WINDOW_HEIGHT-1, COLOR[255], COLOR[0], "20")

		printf_tb(rwidth, rheight, COLOR[120], COLOR[0], "*")
		printf_tb(rwidth, rheight-2, COLOR[120], COLOR[0], "N")
		printf_tb(rwidth, rheight+2, COLOR[120], COLOR[0], "S")
		printf_tb(rwidth+3, rheight, COLOR[120], COLOR[0], "E")
		printf_tb(rwidth-3, rheight, COLOR[120], COLOR[0], "W")
		printf_tb(rwidth-2, rheight-1, COLOR[120], COLOR[0], "NW")
		printf_tb(rwidth+1, rheight-1, COLOR[120], COLOR[0], "NE")
		printf_tb(rwidth-2, rheight+1, COLOR[120], COLOR[0], "SW")
		printf_tb(rwidth+1, rheight+1, COLOR[120], COLOR[0], "SE")

    for i := 0; i < 99; i++ {
      s1 := sections[rand.Intn(len(sections))]
      s2 := sections[rand.Intn(len(sections))]
      igor.Connect(s1[0], s1[1], s2[0], s2[1], MAP)
    }

		printf_tb((igor.WINDOW_WIDTH/2)-8, 0, COLOR[32], COLOR[0], "--- I.G.O.R. ---")
    for x := 0; x < len(MAP.Tiles); x++ {
      row := MAP.Tiles[x]
      for y := 0; y < len(MAP.Tiles[x]); y++ {
        t := row[y]
        printf_tb(t.X, t.Y, COLOR[120], COLOR[0], t.I)
      }
    }
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
