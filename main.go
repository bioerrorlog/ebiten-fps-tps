package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	startTime     time.Time
	lastDebugTime time.Time
	updateCount   int
	drawCount     int
	updateLogged  bool
	drawLogged    bool
}

func (g *Game) Update() error {
	now := time.Now()

	// 最初の10秒間、Update()の呼び出し回数をカウント
	if now.Sub(g.startTime) <= 10*time.Second {
		g.updateCount++
	} else if !g.updateLogged {
		log.Printf("10秒間でUpdate()が呼び出された回数: %d", g.updateCount)
		g.updateLogged = true
	}

	// 1秒ごとにTPSとFPSをデバッグプリント
	if now.Sub(g.lastDebugTime) >= time.Second {
		g.lastDebugTime = now
		log.Printf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	now := time.Now()

	// 最初の10秒間、Draw()の呼び出し回数をカウント
	if now.Sub(g.startTime) <= 10*time.Second {
		g.drawCount++
	} else if !g.drawLogged {
		log.Printf("10秒間でDraw()が呼び出された回数: %d", g.drawCount)
		g.drawLogged = true
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{
		startTime:     time.Now(),
		lastDebugTime: time.Now(),
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("FPSとTPSのデバッグプリント")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
