package main

import (
	"fmt"
	"log"
	"math/rand"
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
	initEntities(&board)
	_, err := boardView.Init(int32(boardX), int32(boardY), 25, "Squirrel")
	if err != nil {
		log.Panic(err)
	}

	if err = boardView.Update(board.GetBoard(), board.player); err != nil {
		fmt.Println(err)
	}
	go HandleEvents(events)

	go func(trigger chan<- bool) {
		step := 0
		for {

			goodPlants := board.getEntities(&GoodPlant{})
			goodPlant := board.getEntityByPath(board.board, board.player, goodPlants, true)
			if goodPlant != nil {
				board.player.setPath(generatePath(board.board, board.player, xy{goodPlant.getX(), goodPlant.getY()}))
				p := board.player.getPath()
				board.move(board.player, xy{p[len(p)-2].x, p[len(p)-2].y})
			}

			// beasts
			for _, beast := range board.beasts {
				if step%4 == 0 {
					target := xy{}
					if beast.getName() == "BadBeast" {
						target.x = board.player.getX()
						target.y = board.player.getY()
					} else { // GoodBeast
						target.x = rand.Intn(xMax)
						target.y = rand.Intn(yMax)
					}
					beast.setPath(generatePath(board.board, beast, target))
					p := beast.getPath()
					if len(p) >= 2 {
						board.move(beast, xy{p[len(p)-2].x, p[len(p)-2].y})
					}
				}
			}
			step++
			time.Sleep(time.Millisecond * 200)
			trigger <- true
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
