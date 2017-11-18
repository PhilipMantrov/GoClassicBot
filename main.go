package main

import (
	"./module/Operation/KeyHandler"
	"./module/Operation/MouseHandler"
	"./module/Operation/ScreenHandler"
)

func main()  {
	KeyHandler.TestKeyBoard();
	MouseHandler.TestMouse();
	ScreenHandler.TestScreen();
}
