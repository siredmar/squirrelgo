package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

var boardView BoardView = BoardView{}

func main() {
	events := make(chan InputEvents)
	updateTrigger := make(chan bool)
	quit := make(chan bool)

	boardX = 25
	boardY = 10

	board := InitBoard(int(boardX), int(boardY))
	board.Create()
	var i int32

	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(int(i), 0), int(i), 0)
		board.AddEntity(createWall(int(i), int(boardY-1)), int(i), int(boardY-1))
	}
	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(0, int(i)), 0, int(i))
		board.AddEntity(createWall(int(boardX-1), int(i)), int(boardX-1), int(i))
	}

	board.AddEntity(createWall(5, 1), 5, 1)
	board.AddEntity(createWall(5, 2), 5, 2)
	board.AddEntity(createWall(5, 3), 5, 3)
	board.AddEntity(createWall(5, 4), 5, 4)
	board.AddEntity(createWall(4, 4), 4, 4)
	board.AddEntity(createWall(3, 4), 3, 4)
	board.AddEntity(createWall(17, 8), 17, 8)
	board.AddEntity(createWall(17, 8), 17, 7)
	board.AddEntity(createWall(17, 6), 17, 6)
	board.AddEntity(createWall(17, 5), 17, 5)
	board.AddEntity(createWall(16, 5), 16, 5)
	board.AddEntity(createWall(15, 5), 15, 5)
	board.AddEntity(createWall(14, 4), 14, 5)
	board.AddEntity(createWall(14, 4), 14, 4)
	board.AddEntity(createGoodPlant(10, 4), 10, 4)
	board.AddEntity(createGoodPlant(4, 3), 4, 3)
	board.AddEntity(createGoodPlant(2, 8), 2, 8)
	board.AddEntity(createGoodPlant(20, 3), 20, 3)
	board.AddEntity(createGoodPlant(23, 8), 23, 8)
	board.AddEntity(createBadPlant(2, 5), 2, 5)
	board.AddEntity(createBadPlant(7, 7), 7, 7)
	board.AddEntity(createBadPlant(8, 1), 8, 1)
	board.AddEntity(createBadPlant(21, 5), 21, 5)
	player := createMasterSquirrel(10, 5)
	board.AddEntity(player, 10, 5)
	board.addPlayer(player)

	_, err := boardView.Init(boardX, boardY, 25, "Squirrel")
	if err != nil {
		log.Panic(err)
	}

	if err = boardView.Update(board.GetBoard()); err != nil {
		fmt.Println(err)
	}

	go HandleEvents(events)

	go func(t <-chan bool) {
		for {
			<-t
			fmt.Println("Update boardview")
			boardView.Update(board.GetBoard())
		}
	}(updateTrigger)

	go func(e chan InputEvents, trigger chan<- bool, q chan<- bool) {
		for {
			ev := <-e
			if ev.quit == true {
				quit <- true
			}
			if ev.up == true {
				y := board.player.getY()
				x := board.player.getX()
				fmt.Println("current: ", x, y, "new: ", x, y-1)
				board.move(board.player, x, y-1)
				board.AddEntity(createNone(x, y), x, y)
				trigger <- true
			}
			if ev.right == true {
				y := board.player.getY()
				x := board.player.getX()
				fmt.Println("current: ", x, y, "new: ", x+1, y)
				board.move(board.player, x+1, y)
				board.AddEntity(createNone(x, y), x, y)
				trigger <- true
			}
			if ev.down == true {
				y := board.player.getY()
				x := board.player.getX()
				fmt.Println("current: ", x, y, "new: ", x, y+1)
				board.move(board.player, x, y+1)
				board.AddEntity(createNone(x, y), x, y)
				trigger <- true
			}
			if ev.left == true {
				y := board.player.getY()
				x := board.player.getX()
				fmt.Println("current: ", x, y, "new: ", x-1, y)
				board.move(board.player, x-1, y)
				board.AddEntity(createNone(x, y), x, y)
				trigger <- true
			}
		}
	}(events, updateTrigger, quit)

	for {
		<-quit
		fmt.Println("exiting")
		os.Exit(0)
	}
	boardView.Cleanup()
}

func HandleEvents(e chan<- InputEvents) {
	var ret InputEvents
	fmt.Println("Started go routine HandleEvnts()")
	for {
		time.Sleep(time.Millisecond * 50)
		event := sdl.PollEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			ret = InputEvents{}
			ret.quit = true
			e <- ret
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE && t.State == 1 {
				ret = InputEvents{}
				ret.quit = true
				fmt.Println("Escape")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_UP && t.State == 1 {
				ret = InputEvents{}
				ret.up = true
				fmt.Println("Up")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_RIGHT && t.State == 1 {
				ret = InputEvents{}
				ret.right = true
				fmt.Println("Right")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_DOWN && t.State == 1 {
				ret = InputEvents{}
				ret.down = true
				fmt.Println("Down")
				e <- ret
			} else if t.Keysym.Sym == sdl.K_LEFT && t.State == 1 {
				ret = InputEvents{}
				ret.left = true
				fmt.Println("Left")
				e <- ret
			}
		}
	}
}
