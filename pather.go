package main

import (
	"fmt"

	astar "github.com/beefsack/go-astar"
)

// Kind* constants refer to tile kinds for input and output.
const (
	KindNone = iota
	KindMasterSquirrel
	KindWall
	KindGoodBeast
	KindBadBeast
	KindGoodPlant
	KindBadPlant
)

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
	return t.W.costs[toT.Kind]
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
type World struct {
	world map[int]map[int]*Tile
	costs map[int]float64
}

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w.world[x] == nil {
		return nil
	}
	return w.world[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w.world[x] == nil {
		w.world[x] = map[int]*Tile{}
	}
	w.world[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

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
func ParseWorld(board [][]Entity, costs map[int]float64) World {
	w := World{}
	w.world = map[int]map[int]*Tile{}
	w.costs = costs
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
