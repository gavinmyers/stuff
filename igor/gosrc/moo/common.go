package moo

//#0
type Thing struct {
  Id int
  Parent *Thing
  Children []*Thing
}

type Area struct {
  Things []*Thing
}

type World struct {
  Things []*Thing
  Areas []*Area
}
