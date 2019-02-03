package main

import (
	"image/png"
	"os"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/salviati/go-tmx/tmx"
)

var clearColor = colornames.Skyblue

var tilemap *tmx.Map
var sprites []*pixel.Sprite

type tile struct {
	mapPos   pixel.Vec
	posCoord pixel.Vec
}

var tiles = []*tile{
	// Tree
	&tile{mapPos: pixel.V(3.0, 0.0), posCoord: pixel.V(3.0, 1.0)}, // tree trunk
	&tile{mapPos: pixel.V(2.0, 1.0), posCoord: pixel.V(2.0, 2.0)}, // top-left tree
	&tile{mapPos: pixel.V(3.0, 1.0), posCoord: pixel.V(3.0, 2.0)}, // top-mid tree
	&tile{mapPos: pixel.V(4.0, 1.0), posCoord: pixel.V(4.0, 2.0)}, // top-right tree

	// Top soil
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(0.0, 0.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(1.0, 0.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(2.0, 0.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(3.0, 0.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(4.0, 0.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(5.0, 0.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 7.0), posCoord: pixel.V(6.0, 0.0)}, // ground

	// Lower soil
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(0.0, -1.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(1.0, -1.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(2.0, -1.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(3.0, -1.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(4.0, -1.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(5.0, -1.0)}, // ground
	&tile{mapPos: pixel.V(2.0, 6.0), posCoord: pixel.V(6.0, -1.0)}, // ground

	// Sign
	&tile{mapPos: pixel.V(1.0, 2.0), posCoord: pixel.V(1.0, 1.0)},

	// Bones
	&tile{mapPos: pixel.V(3.0, 2.0), posCoord: pixel.V(4.0, 1.0)},
	&tile{mapPos: pixel.V(4.0, 2.0), posCoord: pixel.V(5.0, 1.0)},
}

func gameloop(win *pixelgl.Window) {
	tm := tilemap.Tilesets[0]
	w := float64(tm.TileWidth)
	h := float64(tm.TileHeight)
	sprite := loadSprite(tm.Image.Source)

	var iX, iY float64
	var fX = float64(tm.TileWidth)
	var fY = float64(tm.TileHeight)

	for !win.Closed() {
		win.Clear(clearColor)

		for _, coord := range tiles {
			iX = coord.mapPos.X * w
			fX = iX + w
			iY = coord.mapPos.Y * h
			fY = iY + h
			sprite.Set(sprite.Picture(), pixel.R(iX, iY, fX, fY))
			pos := coord.posCoord.ScaledXY(pixel.V(w, h))
			sprite.Draw(win, pixel.IM.Moved(pos.Add(pixel.V(0, h))))
		}
		win.Update()
	}
}

func run() {
	// Create the window with OpenGL
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Tilemaps",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	panicIfErr(err)

	// Initialize art assets (i.e. the tilemap)
	tilemap, err = tmx.ReadFile("gameart2d-desert.tmx")
	panicIfErr(err)

	gameloop(win)
}

func loadSprite(path string) *pixel.Sprite {
	f, err := os.Open(path)
	panicIfErr(err)

	img, err := png.Decode(f)
	panicIfErr(err)

	pd := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pd, pd.Bounds())
}

func main() {
	pixelgl.Run(run)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
