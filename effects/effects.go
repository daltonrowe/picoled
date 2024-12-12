package effects

import (
	"image/color"
	"math"
	"time"

	"picoboard/wave"
)

func breathe(leds *[30]color.RGBA) {
	t := float64(time.Now().UnixMilli()) / 1000.0
	w := wave.Sample(t, 3.0)
	c := uint8(math.Round(255.0 * w))

	for i := range leds {
		leds[i] = color.RGBA{R: c, G: 0, B: 0}
	}
}

func sine(leds *[30]color.RGBA) {
	for i := range leds {
		w := wave.Sample(float64(i), 10.0)
		c := uint8(math.Round(255.0*w) - 50)
		leds[i] = color.RGBA{R: 0, G: c, B: 0}
	}
}

func Fill(effect string, leds *[30]color.RGBA) {
	switch effect {
	case "breathe":
		breathe(leds)
	case "sine":
		sine(leds)
	default:
		breathe(leds)
	}
}
