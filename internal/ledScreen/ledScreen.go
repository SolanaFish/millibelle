package ledScreen

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	brightness = 90
	width      = 32
	height     = 8
	ledCount  = width * height
)

type wsEngine interface {
	Init() error
	Render() error
	Wait() error
	Fini()
	Leds(channel int) []uint32
}

var ws wsEngine

func coordinatesToIndex(x int, y int) int {
	if x%2 == 0 {
		return x * height + y
	}
	return (x + 1) * height - y - 1
}

func rgbToColor(r uint8, g uint8, b uint8) uint32 {
	return uint32(r)<<16 + uint32(g)<<8 + uint32(b)
}

func clear() error {
	for i := 0; i < ledCount; i += 1 {
		ws.Leds(0)[i] = rgbToColor(0,0,0)
	}

	return ws.Render()
}

func Init() {
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCount

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}
	ws = dev

	err = ws.Init()
	if err != nil {
		panic(err)
	}
}

func NextFrame(pixels[32][8][3] uint8) error {
	for i := 0; i < 32; i += 1 {
		for j := 0; j < 8; j += 1 {
			r := pixels[i][j][0]
			g := pixels[i][j][1]
			b := pixels[i][j][2]

			index := coordinatesToIndex(i, j)

			ws.Leds(0)[index] = rgbToColor(r, g, b)
		}
	}
	return ws.Render()
}

func Finish() {
	clear()
	ws.Fini()
}
