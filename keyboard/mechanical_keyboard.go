package keyboard

import "fmt"

type MechanicalKeyboard struct {
	SwitchType string
	Size       string
	OS         string
}

func NewMechanicalKeyboard(switchType, size, os string) Keyboard {
	return &MechanicalKeyboard{
		SwitchType: switchType,
		Size:       size,
		OS:         os,
	}
}

func (m MechanicalKeyboard) Input() {
	fmt.Println("Pressing the key on the mechanical keyboard : ", m.SwitchType, m.Size, m.OS)
}
