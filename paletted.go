package spritesheet8x8x256

import (
	"github.com/reiver/go-sprite8x8"
)

type Paletted struct {
	// Pix holds the pixel data, as palette indices.
	Pix []uint8

	// Palette is the color palette used to resolve the colors.
	Palette Palette

	// Sheet is the name of the sprite sheet.
	Category string
}

func (receiver Paletted) slice(id uint8) (low, high int) {
	low  = int(id) * sprite8x8.ByteSize
	high = low     + sprite8x8.ByteSize

	return low, high
}

func (receiver Paletted) sprite(id uint8) []uint8 {
	low, high := receiver.slice(id)

	p := receiver.Pix[low:high]

	return p
}

func (receiver Paletted) Sprite(id uint8) sprite8x8.Paletted {
	p := receiver.sprite(id)

	return sprite8x8.Paletted{
		Pix:      p,
		Palette:  receiver.Palette,
		Category: receiver.Category,
		ID:       id,
	}
}
