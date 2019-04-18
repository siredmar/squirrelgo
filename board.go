package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"

	astar "github.com/beefsack/go-astar"
)

type xy struct {
	x int
	y int
}

var (
	xMax int
	yMax int
)

type Board struct {
	board  [][]Entity
	player Entity
}

func InitBoard(c xy) Board {
	xMax = c.x
	yMax = c.y
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
			b.board[y][x] = createNone(xy{x, y})
		}
	}
}

func (b *Board) addPlayer(e Entity) {
	b.player = e
}

func (b *Board) AddEntity(e Entity, c xy) (bool, error) {
	if c.x >= xMax || c.x < 0 {
		return false, fmt.Errorf("Incorrect x position")
	}
	if c.y >= yMax || c.y < 0 {
		return false, fmt.Errorf("Incorrect y position")
	}
	b.board[c.y][c.x] = e
	return true, nil
}

func (b *Board) RemoveEntity(c xy) (bool, error) {
	if c.x >= xMax || c.x < 0 {
		return false, fmt.Errorf("Incorrect x position")
	}
	if c.y >= yMax || c.y < 0 {
		return false, fmt.Errorf("Incorrect y position")
	}
	b.board[c.y][c.x] = createNone(xy{c.x, c.y})
	return true, nil
}

func (b Board) GetEntity(c xy) Entity {
	return b.board[c.y][c.x]
}

func (b *Board) move(e Entity, new xy) (bool, error) {
	x := e.getX()
	y := e.getY()
	switch v := b.board[new.y][new.x].(type) {

	default:
		return false, fmt.Errorf("unexpected type %T", v)
	case *None:
		e.move(new.x, new.y)
		b.board[new.y][new.x] = e
		b.board[y][x] = createNone(xy{x, y})
		return true, nil
	case *Wall:
		e.updateEnergy(b.board[new.y][new.x].getEnergy())
		return false, nil
	case *GoodBeast:
		return true, nil
	case *BadBeast:
		return true, nil
	case *GoodPlant:
		e.updateEnergy(b.board[new.y][x].getEnergy())
		e.move(new.x, new.y)
		b.board[new.y][new.x] = e
		b.board[y][x] = createNone(xy{x, y})
		b.spawnEntity("goodplant")
		return true, nil
	case *BadPlant:
		e.updateEnergy(b.board[new.y][new.x].getEnergy())
		e.move(new.x, new.y)
		b.board[new.y][new.x] = e
		b.board[y][x] = createNone(xy{x, y})
		b.spawnEntity("badplant")
		return true, nil
	case *MasterSquirrel:
		b.board[y][x] = createNone(xy{x, y})
		return false, nil
	}
}

func generatePath(b [][]Entity, entity Entity, new xy) []point {
	world := ParseWorld(b, entity.getCosts())
	path, _, found := astar.Path(world.Tile(entity.getX(), entity.getY()), world.Tile(new.x, new.y))
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
				b.AddEntity(createGoodPlant(xy{x, y}), xy{x, y})
				break
			case "badplant":
				b.AddEntity(createBadPlant(xy{x, y}), xy{x, y})
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

func getEntityByAirDistance(e []Entity, x, y int, nearest bool) Entity {
	var s []distance
	if len(e) <= 0 {
		return nil
	}

	for i, entity := range e {
		d := math.Sqrt(float64(x-entity.getX())*float64(x-entity.getX()) + float64(y-entity.getY())*float64(y-entity.getY()))
		s = append(s, distance{d, 0, i})
	}
	if nearest == true {
		sort.Sort(distanceSortUp(s))
	} else {
		sort.Sort(distanceSortDown(s))
	}
	return e[s[0].index]
}

func (b Board) countEntitiesInPath(path []point, entity interface{}) int {
	count := 0
	for _, v := range path {
		if reflect.TypeOf(b.board[v.y][v.x]) == reflect.TypeOf(entity) {
			count++
		}
	}
	return count
}

func (b Board) getEntityByPath(board [][]Entity, sourceEntity Entity, e []Entity, nearest bool) Entity {
	var s []distance
	if len(e) <= 0 {
		return nil
	}

	for i, entity := range e {
		p := generatePath(board, b.player, xy{entity.getX(), entity.getY()})
		count := b.countEntitiesInPath(p, &GoodPlant{})
		s = append(s, distance{float64(len(p)), count, i})
	}

	if nearest == true {
		sort.Sort(distanceSortUp(s))
	} else {
		sort.Sort(distanceSortDown(s))
	}
	return e[s[0].index]
}
