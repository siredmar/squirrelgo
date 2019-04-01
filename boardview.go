package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type BoardView struct {
	winWidth      int32
	winHeight     int32
	winTitle      string
	window        *sdl.Window
	renderer      *sdl.Renderer
	canvas        *sdl.Rect
	tileWidth     int32
	tileHeight    int32
	visibleTilesX int32
	visibleTilesY int32
	textures      map[string]*sdl.Texture
}

func (b *BoardView) Init(width int32, heigth int32, title string) (bool, error) {
	b.tileHeight = 45
	b.tileWidth = 45
	b.winTitle = title
	b.visibleTilesX = width
	b.visibleTilesY = heigth
	b.winWidth = b.visibleTilesX * b.tileWidth
	b.winHeight = b.visibleTilesY * b.tileHeight
	b.textures = map[string]*sdl.Texture{}

	var err error

	sdl.Init(sdl.INIT_VIDEO)

	b.window, err = sdl.CreateWindow(b.winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		b.winWidth, b.winHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		return false, fmt.Errorf("Failed to create window: %s\n", err)
	}

	b.renderer, err = sdl.CreateRenderer(b.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return false, fmt.Errorf("Failed to create renderer: %s\n", err)
	}

	if err = b.loadAllTextures(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return true, nil
}

func (b BoardView) Cleanup() {
	b.window.Destroy()
	b.renderer.Destroy()
	for _, v := range b.textures {
		v.Destroy()
	}
}

func (b BoardView) DrawGrid(enabled bool) {
	if enabled == true {
		b.renderer.SetDrawColor(255, 255, 255, 255)

		var x, y int32
		for y = 0; y < b.winHeight; y = y + b.tileHeight {
			b.renderer.DrawLine(0, y, b.winWidth, y)
			for x = 0; x < b.winWidth; x = x + b.tileWidth {
				b.renderer.DrawLine(x, 0, x, b.winHeight)
			}
		}
		b.renderer.Present()
	} else {
		b.renderer.Clear()
	}
}

func (b BoardView) loadAllTextures() error {
	var t *sdl.Texture

	t, err := b.LoadTexture("gfx/players/player0.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["player0"] = t

	t, err = b.LoadTexture("gfx/goodbeasts/goodbeast0.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["goodbeast0"] = t

	t, err = b.LoadTexture("gfx/badbeasts/badbeast0.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["badbeast0"] = t

	t, err = b.LoadTexture("gfx/badplants/badplant0.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["badplant0"] = t

	t, err = b.LoadTexture("gfx/badplants/badplant1.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["badplant1"] = t

	t, err = b.LoadTexture("gfx/goodplants/goodplant0.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["goodplant0"] = t

	t, err = b.LoadTexture("gfx/goodplants/goodplant1.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["goodplant1"] = t

	t, err = b.LoadTexture("gfx/misc/wall0.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["wall0"] = t

	t, err = b.LoadTexture("gfx/misc/wall1.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["wall1"] = t

	t, err = b.LoadTexture("gfx/misc/none.png")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	b.textures["none"] = t

	return nil
}

func (b BoardView) LoadTexture(file string) (*sdl.Texture, error) {
	image, err := img.Load(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to load PNG: %s\n", err)
	}

	texture, err := b.renderer.CreateTextureFromSurface(image)
	if err != nil {
		return nil, fmt.Errorf("Failed to create texture: %s\n", err)
	}
	return texture, nil
}

func (b BoardView) DrawTile(x, y int32, texture string) (bool, error) {
	if x >= b.visibleTilesX || x < 0 || y >= b.visibleTilesY || y < 0 {
		return false, fmt.Errorf("Error: cannot place tile outside borders: %vx%v", x, y)
	}
	t := b.textures[texture]

	src := sdl.Rect{0, 0, b.tileWidth, b.tileWidth}
	dst := sdl.Rect{x * b.tileWidth, y * b.tileHeight, b.tileWidth, b.tileHeight}

	b.renderer.Copy(t, &src, &dst)
	b.renderer.Present()

	return true, nil
}

func (b BoardView) DrawBoard(board [][]Entity) error {
	var x, y int32
	for y = 0; y < b.visibleTilesY; y++ {
		for x = 0; x < b.visibleTilesX; x++ {
			switch v := board[y][x].(type) {
			default:
				return fmt.Errorf("unexpected type %T", v)
			case *None:
				b.DrawTile(x, y, "none")
			case *MasterSquirrel:
				b.DrawTile(x, y, "player0")
			case *Wall:
				b.DrawTile(x, y, "wall0")
			case *GoodBeast:
				b.DrawTile(x, y, "goodbeast0")
			case *BadBeast:
				b.DrawTile(x, y, "badbeast0")
			case *GoodPlant:
				b.DrawTile(x, y, "goodplant0")
			case *BadPlant:
				b.DrawTile(x, y, "badplant0")
			}
		}
	}
	return nil
}

// func (b BoardView) BoardViewUpdate() (bool, error) {

// }