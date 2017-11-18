package ScreenHandler

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func TestScreen() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)
}
