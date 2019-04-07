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

type InputEvents struct {
	up    bool
	down  bool
	left  bool
	right bool
	quit  bool
}

func main() {
	events := make(chan InputEvents)
	updateTrigger := make(chan bool)

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

	// if err = boardView.DrawBoard(board.GetBoard()); err != nil {
	// 	fmt.Println(err)
	// }

	go HandleEvents(events)
	go boardView.Update(board.GetBoard(), updateTrigger)

	for {
		e := <-events
		if e.quit == true {
			break
		} else {
			updateTrigger <- true
		}
	}
	boardView.Cleanup()

}

func HandleEvents(e chan InputEvents) {
	var ret InputEvents
	fmt.Println("Started go routine HandleEvnts()")
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		fmt.Printf("Event: %T\n", event)
		switch t := event.(type) {
		case *sdl.QuitEvent:
			ret = InputEvents{}
			ret.quit = true
			e <- ret
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				ret = InputEvents{}
				ret.quit = true
				fmt.Println("Escape")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_UP {
				ret = InputEvents{}
				ret.up = true
				fmt.Println("Up")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_RIGHT {
				ret = InputEvents{}
				ret.right = true
				fmt.Println("Right")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_DOWN {
				ret = InputEvents{}
				ret.down = true
				fmt.Println("Down")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_LEFT {
				ret = InputEvents{}
				ret.left = true
				fmt.Println("Left")
				e <- ret
			}
		}
	}
}
