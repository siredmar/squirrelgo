package main

import "fmt"

var (
	xMax int
	yMax int
)

type Board struct {
	board [][]Entity
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
			b.board[y][x] = createNone()
		}
	}
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
	b.board[y][x] = createNone()
	return true, nil
}

func (b Board) GetEntity(x, y int) Entity {
	return b.board[y][x]
}
