package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) resetGameplay() {
	g.grid = [9]Cell{}
	g.turn = HUMANO
	g.gameOver = false
	g.message = ""
}

func (g *Game) updateGameplay() error {
	for i := range g.grid {
		if g.grid[i].value != VAZIO {
			g.grid[i].frameTimer++
			if g.grid[i].frameTimer >= AnimationSpeed {
				g.grid[i].frameTimer = 0
				g.grid[i].currentFrame = (g.grid[i].currentFrame + 1) % TotalFrames
			}
		}
	}
	if g.gameOver {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.resetGameplay()
		}
		return nil
	}

	simpleBoard := g.getSimpleBoard()

	if g.turn == HUMANO {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mx, my := ebiten.CursorPosition()
			col := mx / CellSize
			row := my / CellSize

			if col >= 0 && col < 3 && row >= 0 && row < 3 {
				idx := row*3 + col
				if g.grid[idx].value == VAZIO {
					g.grid[idx].value = HUMANO
					g.grid[idx].currentFrame = 0

					simpleBoard = g.getSimpleBoard()
					if checkWinner(simpleBoard, HUMANO) {
						g.message = "GANHOU"
						g.gameOver = true
					} else if isFull(simpleBoard) {
						g.message = "VELHA"
						g.gameOver = true
					} else {
						g.turn = IA
					}
				}
			}
		}
	} else {
		idx := bestMove(simpleBoard)
		if idx != -1 {
			g.grid[idx].value = IA
			g.grid[idx].currentFrame = 0

			simpleBoard = g.getSimpleBoard()
			if checkWinner(simpleBoard, IA) {
				g.message = "PERDEU"
				g.gameOver = true
			} else if isFull(simpleBoard) {
				g.message = "VELHA"
				g.gameOver = true
			} else {
				g.turn = HUMANO
			}
		}
	}
	return nil
}

func (g *Game) drawGameplay(screen *ebiten.Image) {
	opBg := &ebiten.DrawImageOptions{}
	opBg.Filter = ebiten.FilterNearest
	screen.DrawImage(bgImage, opBg)

	for i, cell := range g.grid {
		if cell.value == VAZIO {
			continue
		}

		cellX := (i % 3) * CellSize
		cellY := (i / 3) * CellSize

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(cellX), float64(cellY))
		op.Filter = ebiten.FilterNearest

		sx := cell.currentFrame * FrameWidth
		rect := image.Rect(sx, 0, sx+FrameWidth, FrameHeight)

		if cell.value == HUMANO {
			screen.DrawImage(xAnim.SubImage(rect).(*ebiten.Image), op)
		} else if cell.value == IA {
			screen.DrawImage(oAnim.SubImage(rect).(*ebiten.Image), op)
		}
	}

	if g.gameOver {
		opFade := &ebiten.DrawImageOptions{}
		opFade.GeoM.Scale(GameWidth, GameHeight)
		opFade.ColorScale.Scale(0, 0, 0, 0.6)
		screen.DrawImage(whiteImg, opFade)

		panelW, panelH := 40.0, 14.0
		panelX := (float64(GameWidth) - panelW) / 2
		panelY := (float64(GameHeight) - panelH) / 2

		opPanel := &ebiten.DrawImageOptions{}
		opPanel.GeoM.Scale(panelW, panelH)
		opPanel.GeoM.Translate(panelX, panelY)
		opPanel.ColorScale.Scale(0.1, 0.1, 0.1, 1)
		screen.DrawImage(whiteImg, opPanel)

		textW := len(g.message) * 6
		textX := int(panelX) + (int(panelW)-textW)/2
		textY := int(panelY)

		ebitenutil.DebugPrintAt(screen, g.message, textX, textY)
	}
}
