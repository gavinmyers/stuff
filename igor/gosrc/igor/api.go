package igor

type Tile struct {
  Z int
  X int
  Y int
  I string
}


type Map struct {
  Z int
  Tiles [][]*Tile
}

