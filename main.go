package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	boardX int32
	boardY int32
	event  sdl.Event
	quit   bool
)

func main() {
	boardX = 25
	boardY = 10

	board := InitBoard(int(boardX), int(boardY))
	board.Create()
	var i int32

	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(), int(i), 0)
		board.AddEntity(createWall(), int(i), int(boardY-1))
	}
	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(), 0, int(i))
		board.AddEntity(createWall(), int(boardX-1), int(i))
	}

	board.AddEntity(createWall(), 5, 1)
	board.AddEntity(createWall(), 5, 2)
	board.AddEntity(createWall(), 5, 3)
	board.AddEntity(createWall(), 5, 4)
	board.AddEntity(createWall(), 4, 4)
	board.AddEntity(createWall(), 3, 4)
	board.AddEntity(createWall(), 17, 8)
	board.AddEntity(createWall(), 17, 7)
	board.AddEntity(createWall(), 17, 6)
	board.AddEntity(createWall(), 17, 5)
	board.AddEntity(createWall(), 16, 5)
	board.AddEntity(createWall(), 15, 5)
	board.AddEntity(createWall(), 14, 5)
	board.AddEntity(createWall(), 14, 4)
	board.AddEntity(createGoodPlant(), 10, 4)
	board.AddEntity(createGoodPlant(), 4, 3)
	board.AddEntity(createGoodPlant(), 2, 8)
	board.AddEntity(createGoodPlant(), 20, 3)
	board.AddEntity(createGoodPlant(), 23, 8)
	board.AddEntity(createBadPlant(), 2, 5)
	board.AddEntity(createBadPlant(), 7, 7)
	board.AddEntity(createBadPlant(), 8, 1)
	board.AddEntity(createBadPlant(), 21, 5)
	board.AddEntity(createMasterSquirrel(), 10, 5)

	boardView := BoardView{}
	_, err := boardView.Init(boardX, boardY, 25, "Squirrel")
	if err != nil {
		log.Panic(err)
	}

	if err = boardView.DrawBoard(board.GetBoard()); err != nil {
		fmt.Println(err)
	}
	boardView.DrawStatusBar("Play334er", 123)
	// boardView.DrawGrid(true)
	for !quit {
		HandleEvents()
	}
	boardView.Cleanup()

}

func HandleEvents() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			quit = true
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				quit = true
			}
		}
	}
}
