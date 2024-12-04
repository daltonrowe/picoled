package effects

import (
	"image/color"
	"time"
	"math"
)

func Fill(effect string, leds *[30]color.RGBA) {
	t := float64(time.Now().UnixMilli()) / 1000.0
	f := 5.0

	w := math.Sin(2 * math.Pi * t / f)
	w = (w + 1) / 2

	var red uint8
	red = uint8(math.Round(255.0 * w))

	for i := range leds {
		leds[i] = color.RGBA{R: red, G: 0, B: 0}
	}
}