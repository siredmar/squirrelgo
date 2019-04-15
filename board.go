package main

import (
	"fmt"
	"math/rand"

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

	switch v := b.board[newy][newx].(type) {

	default:
		return false, fmt.Errorf("unexpected type %T", v)
	case *None:
		b.board[newy][newx] = e
		e.move(newx, newy)
		b.board[newy][newx] = e
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
		b.spawnEntity("goodplant")
		return true, nil
	case *BadPlant:
		e.updateEnergy(b.board[newy][newx].getEnergy())
		e.move(newx, newy)
		b.board[newy][newx] = e
		b.spawnEntity("badplant")
		return true, nil
	case *MasterSquirrel:
		return false, nil
	}
}

func (b *Board) generatePath(e Entity, x, y, xnew, ynew int) {
	world := ParseWorld(b.board)
	path, dist, found := astar.Path(world.Tile(1, 1), world.Tile(22, 5))
	if !found {
		fmt.Println("Could not find a path")
	} else {
		fmt.Println(dist)
		entitypath := []point{}
		for _, p := range path {
			pT := p.(*Tile)
			fmt.Println(pT.X, pT.Y)
			a := point{pT.X, pT.Y}
			entitypath = append(entitypath, a)
		}
		e.setPath(entitypath)
	}
}

func (b *Board) spawnEntity(e string) error {

	for {
		x := rand.Intn(xMax)
		y := rand.Intn(yMax)
		switch v := b.board[y][x].(type) {
		default:
			fmt.Printf("spawning failed, coordinates already used: %v", v)
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
