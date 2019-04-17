package main

import (
	"fmt"
	"math/rand"
	"reflect"

	astar "github.com/beefsack/go-astar"
)

var (
	xMax int
	yMax int
)

type Board struct {
	board  [][]Entity
	player Entity
}

func InitBoard(x, y int) Board {
	xMax = x
	yMax = y
	return Board{}
}

func (b Board) SetBoard(newBoard [][]Entity) {
	b.board = newBoard
}

func (b Board) GetBoard() [][]Entity {
	return b.board
}

func (b *Board) Create() {
	b.board = make([][]Entity, yMax)
	for y := range b.board {
		b.board[y] = make([]Entity, xMax)
	}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			b.board[y][x] = createNone(x, y)
		}
	}
}

func (b *Board) addPlayer(e Entity) {
	b.player = e
}

func (b *Board) AddEntity(e Entity, x, y int) (bool, error) {
	if x >= xMax || x < 0 {
		return false, fmt.Errorf("Incorrect x position")
	}
	if y >= yMax || y < 0 {
		return false, fmt.Errorf("Incorrect y position")
	}
	b.board[y][x] = e
	return true, nil
}

func (b *Board) RemoveEntity(x, y int) (bool, error) {
	if x >= xMax || x < 0 {
		return false, fmt.Errorf("Incorrect x position")
	}
	if y >= yMax || y < 0 {
		return false, fmt.Errorf("Incorrect y position")
	}
	b.board[y][x] = createNone(x, y)
	return true, nil
}

func (b Board) GetEntity(x, y int) Entity {
	return b.board[y][x]
}

func (b *Board) move(e Entity, newx, newy int) (bool, error) {
	x := e.getX()
	y := e.getY()
	switch v := b.board[newy][newx].(type) {

	default:
		return false, fmt.Errorf("unexpected type %T", v)
	case *None:
		e.move(newx, newy)
		b.board[newy][newx] = e
		b.board[y][x] = createNone(x, y)
		return true, nil
	case *Wall:
		e.updateEnergy(b.board[newy][newx].getEnergy())
		return false, nil
	case *GoodBeast:
		return true, nil
	case *BadBeast:
		return true, nil
	case *GoodPlant:
		e.updateEnergy(b.board[newy][newx].getEnergy())
		e.move(newx, newy)
		b.board[newy][newx] = e
		b.board[y][x] = createNone(x, y)
		b.spawnEntity("goodplant")
		return true, nil
	case *BadPlant:
		e.updateEnergy(b.board[newy][newx].getEnergy())
		e.move(newx, newy)
		b.board[newy][newx] = e
		b.board[y][x] = createNone(x, y)
		b.spawnEntity("badplant")
		return true, nil
	case *MasterSquirrel:
		b.board[y][x] = createNone(x, y)
		return false, nil
	}
}

func (b *Board) generatePath(x, y, xnew, ynew int) []point {
	world := ParseWorld(b.board)
	path, _, found := astar.Path(world.Tile(x, y), world.Tile(xnew, ynew))
	if !found {
		fmt.Println("Could not find a path")
	} else {
		entitypath := []point{}
		for _, p := range path {
			pT := p.(*Tile)
			a := point{pT.X, pT.Y}
			entitypath = append(entitypath, a)
		}
		return entitypath
	}
	return nil
}

func (b *Board) spawnEntity(e string) error {

	for {
		x := rand.Intn(xMax)
		y := rand.Intn(yMax)
		switch b.board[y][x].(type) {
		default:
			continue
		case *None:
			switch e {
			default:
				return fmt.Errorf("unknown entity type to spawn")
			case "goodplant":
				b.AddEntity(createGoodPlant(x, y), x, y)
				break
			case "badplant":
				b.AddEntity(createBadPlant(x, y), x, y)
				break
			}
		}
		break
	}
	return nil
}

func (b Board) getEntities(v interface{}) []Entity {
	var x, y int
	var entities []Entity
	for y = 0; y < yMax; y++ {
		for x = 0; x < xMax; x++ {
			if reflect.TypeOf(b.board[y][x]) == reflect.TypeOf(v) {
				entities = append(entities, b.board[y][x])
			}
		}
	}
	return entities
}
