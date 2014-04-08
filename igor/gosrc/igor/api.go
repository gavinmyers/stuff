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
    if(i < lenX) {
      if(startX > startY) {
        path[i] = DirWest
      } else {
        path[i] = DirEast
      }
    } else {
      if(startY > endY) {
        path[i] = DirNorth
      } else {
        path[i] = DirSouth
      }
    }
	}
  perm := rand.Perm(len(path))
  tiles := make([]*Tile, lenT)
  pathX := startX
  pathY := startY
  for i, v := range perm {
    if path[v] == DirWest {
      pathX--
    } else if path[v] == DirEast {
      pathX++
    } else if path[v] == DirNorth {
      pathY--
    } else if path[v] == DirSouth {
      pathY++
    }
    t := &Tile {X:pathX, Y:pathY, I:"."}
    tiles[i] = t
    if(pathX > 0 && pathY > 0) {
      m.Tiles[pathX][pathY] = t
    }
  }
}

