package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"

	"picoboard/effects"
)

var (
	leds   [42]color.RGBA
	effect string
)

func main() {
	effect = "red-green-blue"

	pin := machine.GPIO22
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.NewWS2812(pin)

	for {
		effects.Fill(effect, &leds)
		ws.WriteColors(leds[:])
		time.Sleep(16 * time.Millisecond)
	}
}
