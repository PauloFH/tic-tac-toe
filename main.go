package main

import (
	"log"

	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	GameWidth      = 48
	GameHeight     = 48
	CellSize       = 16
	WindowWidth    = 480
	WindowHeight   = 480
	FrameWidth     = 16
	FrameHeight    = 16
	TotalFrames    = 2
	AnimationSpeed = 15
)

const (
	HUMANO = "X"
	IA     = "O"
	VAZIO  = ""
)

type GameState int

const (
	StateTitle GameState = iota
	StatePlaying
)

type Cell struct {
	value        string
	frameTimer   int
	currentFrame int
}

type Game struct {
	state    GameState
	grid     [9]Cell
	turn     string
	gameOver bool
	message  string
}

func (g *Game) getSimpleBoard() [9]string {
	var b [9]string
	for i, cell := range g.grid {
		b[i] = cell.value
	}
	return b
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.state == StatePlaying {
			g.state = StateTitle
		} else {
			os.Exit(0)
		}
	}
	switch g.state {
	case StateTitle:
		return g.updateTitle()
	case StatePlaying:
		return g.updateGameplay()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateTitle:
		g.drawTitle(screen)
	case StatePlaying:
		g.drawGameplay(screen)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return GameWidth, GameHeight
}

func main() {
	loadAssets()

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Tic-Tac-Toe Pixel Art")
	game := &Game{
		state: StateTitle,
		turn:  HUMANO,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
