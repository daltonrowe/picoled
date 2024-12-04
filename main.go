package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"

	"picoboard/effects"
)

const fps = false

var (
	leds   [30]color.RGBA
	effect string
)

func main() {
	effect = "red-green-blue"

	pin := machine.GPIO22
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.NewWS2812(pin)

	t := time.Now()

	for {
		effects.Fill(effect, &leds)
		ws.WriteColors(leds[:])
		time.Sleep(16 * time.Millisecond)

		if fps {
			s := time.Since(t)
			println(1000/s.Milliseconds(), " fps")
			t = time.Now()
		}
	}
}
