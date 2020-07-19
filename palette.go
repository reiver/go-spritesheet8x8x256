package spritesheet8x8x256

import (
	"image/color"
)

type Palette interface {
	color.Model
	Color(uint8) color.Color
}
