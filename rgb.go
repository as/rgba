// Package rgba provides a conversion between 32-bit RGBA quad values
// and the standard library's image and color packages. This package
// defines no types of its own and is used primarily to translate from
// a common web notation to the standard library
package rgba

import (
	"image"
	"image/color"
	"image/color/palette"
	"math/rand"
	"runtime"
)

// Hex converts a 32-bit RGBA quad to a color.RGBA
func Hex(rgba uint32) color.RGBA {
	return color.RGBA{
		R: uint8(rgba >> 24),
		G: uint8(rgba << 8 >> 24),
		B: uint8(rgba << 16 >> 24),
		A: uint8(rgba << 24 >> 24),
	}
}

func Plan9(c color.Color) color.Color {
	return palette.Plan9[color.Palette(palette.Plan9).Index(c)]
}

// Uniform is short for image.NewUniform(Hex(rgba)). On linux,
// we are doing something nasty by pre-swizzling the uniform
// colors. This is until I can fix the swizzle in as/shiny for linux
var Uniform = func() func(rgba uint32) *image.Uniform {
	if runtime.GOOS == "linux" {
		return linuxuniform
	}
	return uniform
}()

// Uint32 converts a color.RGBA to a uint32
func Uint32(c color.RGBA) uint32 {
	return uint32(c.R)<<24 | uint32(c.G)<<16 | uint32(c.B)<<8 | uint32(c.A)
}

func Rand() *image.Uniform {
	c := palette.Plan9[rand.Intn(len(palette.Plan9))].(color.RGBA)
	return image.NewUniform(c)
}

// uniform is short for image.NewUniform(Hex(rgba))
func uniform(rgba uint32) *image.Uniform {
	return image.NewUniform(Hex(rgba))
}

// linuxuniform is short for image.NewUniform(Hex(rgba))\
// this function exists because I haven't fixed swizzle on
// linux yet.
func linuxuniform(rgba uint32) *image.Uniform {
	c := Hex(rgba)
	c.R, c.B = c.B, c.R
	return image.NewUniform(c)
}
