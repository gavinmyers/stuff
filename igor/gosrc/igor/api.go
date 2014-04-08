package igor
import "math/rand"

func Connect(startX, startY, endX, endY int, m *Map) {
	lenX := startX - endX
	if lenX < 0 {
		lenX = endX - startX
	}
	lenY := startY - endY
	if lenY < 0 {
		lenY = endY - startY
	}
	lenT := lenX + lenY
	path := make([]int, lenT, lenT)
	for i := 0; i < lenT; i++ {
    if(startY > endY) {
      path[i] = DirNorth
    } else {
      path[i] = DirSouth
    }
  }
	for i := 0; i < lenX; i++ {
    if(startX > startY) {
      path[i] = DirWest
    } else {
      path[i] = DirEast
    }
	}
  dest := make([]int, len(path))
  perm := rand.Perm(len(path))
  for i, v := range perm {
    dest[v] = path[i]
  }
  tiles := make([]*Tile, lenT)
  pathX := startX
  pathY := startY
  for i := 0; i < len(dest); i++ {
    if dest[i] == DirWest {
      pathX--
    } else if dest[i] == DirEast {
      pathX++
    } else if dest[i] == DirNorth {
      pathY--
    } else if dest[i] == DirSouth {
      pathY++
    }
    t := &Tile {X:pathX, Y:pathY, I:"."}
    tiles[i] = t
    if(pathX > 0 && pathY > 0) {
      m.Tiles[pathX][pathY] = t
    }
  }
}

