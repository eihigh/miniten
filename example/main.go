package main

import (
	"embed"
	"math/rand/v2"
	"time"

	"github.com/eihigh/miniten"
)

//go:embed *.png
var fsys embed.FS

var (
	x     = 0
	y     = 100
	t     = 0
	start = time.Now()
)

func main() {
	miniten.Run(draw)
}

func draw() {
	t++
	miniten.Println(t, time.Since(start))
	miniten.DrawRect(x, y, 100, 100)
	miniten.Println("Hello,", "World!")
	miniten.Println("こんにちは、世界")
	miniten.Println(42)
	miniten.Println("クリックで箱を移動")
	if miniten.IsClicked() {
		x, y = miniten.CursorPos()
	}
	miniten.DrawImage("smile.png", rand.N(15), rand.N(15)+200)
	miniten.DrawImageFS(fsys, "smile.png", rand.N(15), rand.N(15)+300)
}
