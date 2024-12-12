package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"

	"picoboard/effects"
	"picoboard/serial"
)

const fps = false

var (
	leds   [30]color.RGBA
	effect string
)

func main() {
	effect = "roll"

	pin := machine.GPIO22
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.NewWS2812(pin)

	t := time.Now()

	for {
		new, err := serial.Read()

		if err == nil {
			effect = new
			println(effect)
		}

		effects.Fill(effect, &leds)
		ws.WriteColors(leds[:])

		r := time.Duration(20 - time.Since(t).Milliseconds())
		time.Sleep(r * time.Millisecond)

		if fps {
			s := time.Since(t)
			println(1000/s.Milliseconds(), " fps")
			t = time.Now()
		}
	}
}
