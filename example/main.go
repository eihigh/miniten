package main

import (
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

var (
	x = 0
	y = 100
)

func main() {
	miniten.Run(draw)
}

func draw() {
	miniten.DrawRect(x, y, 100, 100)
	miniten.Println("Hello,", "World!")
	miniten.Println("こんにちは、世界")
	miniten.Println(42)
	miniten.Println("クリックで箱を移動")
	if miniten.IsClicked() {
		x, y = miniten.CursorPos()
	}
	miniten.DrawImage("smile.png", rand.N(15), rand.N(15)+200)
}
