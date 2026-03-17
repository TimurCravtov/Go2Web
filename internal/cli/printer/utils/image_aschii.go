package utils

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/image/draw"
)

func ImageToAscii(imageData []byte, width int, height int) (string, error) {
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return "", fmt.Errorf("decode error: %w", err)
	}

	// Resize to target dimensions
	destRect := image.Rect(0, 0, width, height)
	dest := image.NewRGBA(destRect)
	draw.BiLinear.Scale(dest, dest.Bounds(), img, img.Bounds(), draw.Over, nil)

	// Density map (ordered from dark/empty to light/dense)
	// Swap the order if your terminal background is light
	asciiChars := []rune{' ', '.', ':', '-', '=', '+', '*', '#', '%', '@'}
	var sb strings.Builder

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := dest.At(x, y)
			r16, g16, b16, _ := c.RGBA()

			// Convert 16-bit color (0-65535) to 8-bit (0-255)
			r := uint8(r16 >> 8)
			g := uint8(g16 >> 8)
			b := uint8(b16 >> 8)

			// Calculate luminance for character selection
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			idx := int(lum * float64(len(asciiChars)-1) / 255.0)

			// Generate the colored character string
			// We add a trailing space to maintain a roughly square aspect ratio in terminals
			charStr := color.RGB(r, g, b).Sprintf("%c ", asciiChars[idx])
			sb.WriteString(charStr)
		}
		sb.WriteByte('\n')
	}

	return sb.String(), nil
}