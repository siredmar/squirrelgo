package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	boardX int
	boardY int
	quit   bool
)

var boardView BoardView = BoardView{}

func update(board Board, updateTrigger chan bool) int {

	sdl.Do(func() {
		for {
			<-updateTrigger
			boardView.Update(board.GetBoard(), board.player)
		}
	})

	return 1
}

func main() {
	events := make(chan InputEvents)
	updateTrigger := make(chan bool)

	boardX = 25
	boardY = 10

	board := InitBoard(xy{boardX, boardY})
	board.Create()
	var i int

	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(xy{i, 0}), xy{i, 0})
		board.AddEntity(createWall(xy{i, boardY - 1}), xy{i, boardY - 1})
	}
	for i = 0; i < boardX; i++ {
		board.AddEntity(createWall(xy{0, i}), xy{0, i})
		board.AddEntity(createWall(xy{boardX - 1, i}), xy{boardX - 1, i})
	}

	board.AddEntity(createWall(xy{5, 1}), xy{5, 1})
	board.AddEntity(createWall(xy{5, 2}), xy{5, 2})
	board.AddEntity(createWall(xy{5, 3}), xy{5, 3})
	board.AddEntity(createWall(xy{5, 4}), xy{5, 4})
	board.AddEntity(createWall(xy{4, 4}), xy{4, 4})
	board.AddEntity(createWall(xy{3, 4}), xy{3, 4})
	board.AddEntity(createWall(xy{17, 8}), xy{17, 8})
	board.AddEntity(createWall(xy{17, 8}), xy{17, 7})
	board.AddEntity(createWall(xy{17, 6}), xy{17, 6})
	board.AddEntity(createWall(xy{17, 5}), xy{17, 5})
	board.AddEntity(createWall(xy{16, 5}), xy{16, 5})
	board.AddEntity(createWall(xy{15, 5}), xy{15, 5})
	board.AddEntity(createWall(xy{14, 4}), xy{14, 5})
	board.AddEntity(createWall(xy{14, 4}), xy{14, 4})
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("goodplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")
	board.spawnEntity("badplant")

	player := createMasterSquirrel(xy{10, 5})
	board.AddEntity(player, xy{10, 5})
	board.addPlayer(player)

	_, err := boardView.Init(int32(boardX), int32(boardY), 25, "Squirrel")
	if err != nil {
		log.Panic(err)
	}

	if err = boardView.Update(board.GetBoard(), board.player); err != nil {
		fmt.Println(err)
	}
	go HandleEvents(events)

	go func(trigger chan<- bool) {
		for {
			entities := board.getEntities(&GoodPlant{})
			a := board.getEntityByPath(board.board, board.player, entities, true)
			if a != nil {
				xnew := a.getX()
				ynew := a.getY()
				board.player.setPath(generatePath(board.board, board.player, xy{xnew, ynew}))
				p := board.player.getPath()
				board.move(board.player, xy{p[len(p)-2].x, p[len(p)-2].y})
				time.Sleep(time.Millisecond * 200)
				trigger <- true
			}
		}
	}(updateTrigger)

	go func(e chan InputEvents, trigger chan<- bool) {
		for {
			ev := <-e
			if ev.quit == true {
				os.Exit(0)
				break
			}
			// manual controlling the player

			// y := board.player.getY()
			// x := board.player.getX()
			// var s bool
			// newx := x
			// newy := y

			// if ev.up == true {
			// 	newy += -1
			// } else if ev.right == true {
			// 	newx += 1
			// } else if ev.down == true {
			// 	newy += 1
			// } else if ev.left == true {
			// 	newx += -1
			// }

			// s, _ = board.move(board.player, newx, newy)
			// entities := board.getEntities(&GoodPlant{})
			// a := getEntityByAirDistance(entities, board.player.getX(), board.player.getY(), true)

			// fmt.Println(a)
			// board.player.setPath(generatePath(board.board, board.player, a.getX(), a.getY()))
			// if s == true {
			// 	board.AddEntity(createNone(x, y), x, y)
			// }
			// trigger <- true
		}
	}(events, updateTrigger)

	var exitcode int
	sdl.Main(func() {
		exitcode = update(board, updateTrigger)
	})

	os.Exit(exitcode)
}
