package ledScreen

import (
	"fmt"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	_ "image/png"
)

const (
	brightness = 90
	width      = 32
	height     = 8
	fps		   = 10
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
var running = false

func coordinatesToIndex(x int, y int) int {
	//fmt.Println(x, y, x * height + y)
	return x * height + y
}

func rgbToColor(r uint8, g uint8, b uint8) uint32 {
	return uint32(r)<<16 + uint32(g)<<8 + uint32(b)
}

//func (inv *invader) display() error {
//	bounds := inv.img[inv.current].Bounds()
//	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
//		for x := bounds.Min.X; x < bounds.Max.X; x++ {
//			r, g, b, _ := inv.img[inv.current].At(x, y).RGBA()
//			inv.ws.Leds(0)[coordinatesToIndex(bounds, x, y)] = rgbToColor(r, g, b)
//		}
//	}
//	return inv.ws.Render()
//}

func clear() error {
	for i := 0; i < ledCount; i += 1 {
		ws.Leds(0)[i] = rgbToColor(0,255,0)
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

	//for count := 0; count < 50; count++ {
	//	inv.display()
	//	inv.next()
	//	time.Sleep(200 * time.Millisecond)
	//}
}

func NextFrame(pixels[32][8][3] uint8) error {
	for i := 0; i < 32; i += 1 {
		for j := 0; j < 8; j += 1 {
			r := pixels[i][j][0]
			g := pixels[i][j][1]
			b := pixels[i][j][2]

			index := coordinatesToIndex(i, j)

			rgb := rgbToColor(r, g, b)

			fmt.Println(index, r, g, b, rgb)

			ws.Leds(0)[index] = rgb
		}
	}
	return ws.Render()
}

func Finish() {
	clear()
	ws.Fini()
}
