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

func roll(leds *[30]color.RGBA) {
	t := float64(time.Now().UnixMilli()) / 1000.0
	o := wave.Sample(t, 6.0) * 100.0

	for i := range leds {
		w := wave.Sample(float64(i)+o, 10.0)
		c := uint8(math.Round(255.0*w) - 50)
		leds[i] = color.RGBA{R: 0, G: 0, B: c}
	}
}

func multiroll(leds *[30]color.RGBA) {
	t := float64(time.Now().UnixMilli()) / 1000.0
	rof := wave.Sample(t, 40.0) * 100.0
	gof := wave.Sample(t, 30.0) * 100.0
	bof := wave.Sample(t, 20.0) * 100.0

	for i := range leds {
		rw := wave.Sample(float64(i+int(rof)), 10.0)
		gw := wave.Sample(float64(i+int(gof)), 10.0)
		bw := wave.Sample(float64(i+int(bof)), 10.0)

		r := uint8(math.Round(255.0*rw) - 50)
		g := uint8(math.Round(255.0*gw) - 50)
		b := uint8(math.Round(255.0*bw) - 50)

		leds[i] = color.RGBA{R: r, G: g, B: b}
	}
}

func Fill(effect string, leds *[30]color.RGBA) {
	switch effect {
	case "breathe":
		breathe(leds)
	case "sine":
		sine(leds)
	case "roll":
		roll(leds)
	case "multiroll":
		multiroll(leds)
	default:
		breathe(leds)
	}
}
