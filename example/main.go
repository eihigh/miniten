package main

import (
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

var x = 0

func main() {
	miniten.Run(draw)
}

func draw() {
	miniten.Println("Hello,", "World!")
	miniten.Println("こんにちは、世界")
	miniten.Println(42)
	miniten.Println("クリックで箱を移動")
	if miniten.IsClicked() {
		x++
	}
	miniten.DrawRect(x, 100, 100, 100)
	miniten.DrawImage("smile.png", rand.N(15), rand.N(15)+200)
}
