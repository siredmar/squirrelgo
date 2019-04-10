package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var events sdl.Event

type InputEvents struct {
	up    bool
	down  bool
	left  bool
	right bool
	quit  bool
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
