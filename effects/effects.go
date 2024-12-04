package effects

import (
	"image/color"
	"time"
)

func Fill(weffect string, leds *[42]color.RGBA) {
	t := time.Now().UnixMilli()


	var red uint8
	var green uint8

	red = 0x00;
	green = 0xff;

	if t % 2 == 0 {
		red = 0xff
		green = 0x00
	}

	for i := range leds {
		leds[i] = color.RGBA{R: red, G: green, B: 0x00}
	}
}