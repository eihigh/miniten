package miniten

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//go:embed Mplus2-Medium.ttf
var font []byte

var (
	draw   func()
	screen *ebiten.Image
	face   *text.GoTextFace
	ln     = 0
	images = map[string]*ebiten.Image{}
)

func Run(d func()) error {
	draw = d

	src, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		panic(err)
	}
	face = &text.GoTextFace{Source: src, Size: 16}

	ebiten.SetWindowSize(1600, 900)
	return ebiten.RunGame(app{})
}

type app struct{}

func (app) Update() error {
	return nil
}

func (app) Draw(s *ebiten.Image) {
	screen = s
	screen.Fill(color.White)
	ln = 0
	draw()
}

func (app) Layout(_, _ int) (int, int) {
	return 1600, 900
}

func IsClicked() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func Println(args ...any) {
	s := fmt.Sprintln(args...)
	opt := &text.DrawOptions{}
	opt.ColorScale.Scale(0, 0, 0, 1)
	for _, line := range strings.Split(s, "\n") {
		opt.GeoM.Reset()
		opt.GeoM.Translate(0, float64(ln)*12)
		text.Draw(screen, line, face, opt)
		ln++
	}
}

func DrawRect(x, y, w, h int) {
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), color.Black, false)
}

func DrawCircle(x, y, r int) {
	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(r), color.Black, false)
}

func DrawImage(path string, x, y int) {
	if _, ok := images[path]; !ok {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if errors.Is(err, os.ErrNotExist) {
			log.Println("画像ファイルが存在しません:", path)
		} else if err != nil {
			log.Println("画像ファイルの読み込みに失敗しました:", err.Error())
		}
		images[path] = img
	}
	img := images[path]
	if img != nil {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(img, opt)
	}
}
