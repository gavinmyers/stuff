package moo

//#0
type Thing struct {
  Id int
  Parent *Thing
  World *World
  Area *Area
  Map *Map
  Tile *Tile
  Children []*Thing
  Sprite Sprite
}

type Sprite struct {
  Id int
  Code string
}


type World struct {
  Areas []*Area
  Maps []*Map
  Things []*Thing
}

type Terrain int32
const (
   Ocean Terrain = 1
   Forest Terrain = 2
   Swamp Terrain = 3
   Plains Terrain = 4
   Tundra Terrain = 5
   Mountain Terrain = 6
   River Terrain = 7
   Lake Terrain = 8
)

type Area struct {
  World *World
  Map *Map
  Things []*Thing
  Terrain Terrain
}

type Map struct {
  World *World
  Area *Area
  Things []*Thing
  Tiles [][]*Tile
}

type Tile struct {
  World *World
  Area *Area
  Map *Map
  Z int
  X int
  Y int
  Things []*Thing
}


