package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed sprites/board.png
var boardPng []byte

//go:embed sprites/X.png
var xPng []byte

//go:embed sprites/O.png
var oPng []byte

var (
	bgImage  *ebiten.Image
	xAnim    *ebiten.Image
	oAnim    *ebiten.Image
	whiteImg *ebiten.Image
)

func loadAssets() {
	img, _, err := image.Decode(bytes.NewReader(boardPng))
	if err != nil {
		log.Fatal("Erro board:", err)
	}
	bgImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(xPng))
	if err != nil {
		log.Fatal("Erro X:", err)
	}
	xAnim = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(oPng))
	if err != nil {
		log.Fatal("Erro O:", err)
	}
	oAnim = ebiten.NewImageFromImage(img)

	whiteImg = ebiten.NewImage(1, 1)
	whiteImg.Fill(color.White)
}
