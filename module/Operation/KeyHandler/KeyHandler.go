package KeyHandler

import (
"fmt"
"github.com/go-vgo/robotgo"
)

func TestKeyBoard() {
	robotgo.TypeString("Hello World")
	robotgo.KeyTap("enter")
	robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")
	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)

	robotgo.WriteAll("Test")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}
}