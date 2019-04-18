package main

// pather_test.go implements a basic world and tiles that implement Pather for
// the sake of testing.  This functionality forms the back end for
// path_test.go, and serves as an example for how to use A* for a grid.

import (
	"fmt"

	astar "github.com/beefsack/go-astar"
)

// Kind* constants refer to tile kinds for input and output.
const (
	// KindPlain (.) is a plain tile with a movement cost of 1.
	KindNone = iota
	KindMasterSquirrel
	KindWall
	KindGoodBeast
	KindBadBeast
	KindGoodPlant
	KindBadPlant
)

// KindCosts map tile kinds to movement costs.
var KindCosts = map[int]float64{
	KindNone:           200.0,
	KindMasterSquirrel: 2000.0,
	KindWall:           100000.0,
	KindGoodBeast:      1000.0,
	KindBadBeast:       2000.0,
	KindGoodPlant:      1.0,
	KindBadPlant:       5000.0,
}

// A Tile is a tile in a grid which implements Pather.
type Tile struct {
	// Kind is the kind of tile, potentially affecting movement.
	Kind int
	// X and Y are the coordinates of the tile.
	X, Y int
	// W is a reference to the World that the tile is a part of.
	W World
}

// PathNeighbors returns the neighbors of the tile, excluding blockers and
// tiles off the edge of the board.
func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil &&
			n.Kind != KindWall {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return KindCosts[toT.Kind]
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

// World is a two dimensional map of Tiles.
type World map[int]map[int]*Tile

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

// FirstOfKind gets the first tile on the board of a kind, used to get the from
// and to tiles as there should only be one of each.
func (w World) FirstOfKind(kind int) *Tile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}
	return nil
}

// // From gets the from tile from the world.
// func (w World) From() *Tile {
// 	return w.FirstOfKind(KindFrom)
// }

// // To gets the to tile from the world.
// func (w World) To() *Tile {
// 	return w.FirstOfKind(KindTo)
// }

// // RenderPath renders a path on top of a world.
// func (w World) RenderPath(path []astar.Pather) string {
// 	width := len(w)
// 	if width == 0 {
// 		return ""
// 	}
// 	height := len(w[0])
// 	pathLocs := map[string]bool{}
// 	for _, p := range path {
// 		pT := p.(*Tile)
// 		pathLocs[fmt.Sprintf("%d,%d", pT.X, pT.Y)] = true
// 	}
// 	rows := make([]string, height)
// 	for x := 0; x < width; x++ {
// 		for y := 0; y < height; y++ {
// 			t := w.Tile(x, y)
// 			r := ' '
// 			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
// 				r = KindRunes[KindPath]
// 			} else if t != nil {
// 				r = KindRunes[t.Kind]
// 			}
// 			rows[y] += string(r)
// 		}
// 	}
// 	return strings.Join(rows, "\n")
// }

func getTileType(e Entity) (int, error) {
	switch v := e.(type) {
	default:
		return -1, fmt.Errorf("unexpected type %T", v)
	case *None:
		return 0, nil
	case *MasterSquirrel:
		return 1, nil
	case *Wall:
		return 2, nil
	case *GoodBeast:
		return 3, nil
	case *BadBeast:
		return 4, nil
	case *GoodPlant:
		return 5, nil
	case *BadPlant:
		return 6, nil
	}
}

// ParseWorld parses a textual representation of a world into a world map.
func ParseWorld(board [][]Entity) World {
	w := World{}
	for y, row := range board {
		for x, entity := range row {
			kind, _ := getTileType(entity)
			w.SetTile(&Tile{
				Kind: kind,
			}, x, y)
		}
	}
	return w
}
