package keyboard

import "fmt"

type NormalKeyboard struct {
	IsCanPress bool
}

func NewNormalKeyboard(isCanPress bool) Keyboard {
	return &NormalKeyboard{
		IsCanPress: isCanPress,
	}
}

func (n NormalKeyboard) Input() {
	fmt.Println("Pressing the key on the keyboard : ", n.IsCanPress)
}
