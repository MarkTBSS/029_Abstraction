package main

import "github.com/MarkTBSS/029_Abstraction/keyboard"

func main() {
	razor := keyboard.NewMechanicalKeyboard("Colourful", "Big", "MacOS")
	razor.Input()

	logitech := keyboard.NewNormalKeyboard(false)
	logitech.Input()
}
