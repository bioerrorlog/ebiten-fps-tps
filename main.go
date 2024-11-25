package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TPS = 60
)

type Game struct {
	updateCount int
	drawCount   int
	perSec      time.Time
}

func (g *Game) Update() error {
	now := time.Now()
	g.updateCount++

	// Debug print per sec
	if now.Sub(g.perSec) >= time.Second {
		log.Printf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
		log.Printf("Update() was called in this sec: %d times", g.updateCount)
		log.Printf("Draw() was called in this sec: %d times\n\n", g.drawCount)

		g.updateCount = 0
		g.drawCount = 0
		g.perSec = now
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawCount++
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{
		perSec: time.Now(),
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Show FPS and TPS with Update()/Draw()")
	ebiten.SetTPS(TPS)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
