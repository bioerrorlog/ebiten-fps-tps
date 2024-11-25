package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	updateCount int
	drawCount   int
	per1Sec     time.Time
	per10Sec    time.Time
}

func (g *Game) Update() error {
	now := time.Now()
	g.updateCount++

	// Debug print TPS and FPS per sec
	if now.Sub(g.per1Sec) >= time.Second {
		log.Printf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
		g.per1Sec = now
	}

	// Debug print Update() and Draw() called count in this 10 sec
	if now.Sub(g.per10Sec) >= 10*time.Second {
		log.Printf("Update() was called in this 10 sec: %d times", g.updateCount)
		log.Printf("Draw() was called in this 10 sec: %d times", g.drawCount)

		g.updateCount = 0
		g.drawCount = 0
		g.per10Sec = now
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
		per1Sec:  time.Now(),
		per10Sec: time.Now(),
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Show FPS and TPS")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
