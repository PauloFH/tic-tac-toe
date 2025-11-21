package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) updateTitle() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if mx >= 4 && mx <= 44 && my >= 20 && my <= 30 {
			g.resetGameplay()
			g.state = StatePlaying
		}
		if mx >= 4 && mx <= 44 && my >= 35 && my <= 45 {
			os.Exit(0)
		}
	}
	return nil
}

func (g *Game) drawTitle(screen *ebiten.Image) {
	opBg := &ebiten.DrawImageOptions{}
	opBg.ColorScale.Scale(0.5, 0.5, 0.5, 1)
	opBg.Filter = ebiten.FilterNearest
	screen.DrawImage(bgImage, opBg)
	ebitenutil.DebugPrintAt(screen, "VELHA", 9, 5)
	opBtn1 := &ebiten.DrawImageOptions{}
	opBtn1.GeoM.Scale(40, 12)
	opBtn1.GeoM.Translate(4, 20)
	opBtn1.ColorScale.Scale(0, 0.5, 0, 1)
	screen.DrawImage(whiteImg, opBtn1)
	ebitenutil.DebugPrintAt(screen, "JOGAR", 9, 18)
	opBtn2 := &ebiten.DrawImageOptions{}
	opBtn2.GeoM.Scale(40, 12)
	opBtn2.GeoM.Translate(4, 35)
	opBtn2.ColorScale.Scale(0.5, 0, 0, 1)
	screen.DrawImage(whiteImg, opBtn2)
	ebitenutil.DebugPrintAt(screen, "SAIR", 12, 33)
}
