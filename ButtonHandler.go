package Antares_vpn

import "github.com/go-playground/colors"

type EventListener interface {
	pressed() bool
	holding() bool
	unpressed() bool
	moving() bool
}

type Button struct {
	text  string
	dims  map[string]string
	color colors.HEXColor
	event EventListener
}

func btnActionPressed()   {}
func btnActionUnPressed() {}
func btnActionHolding()   {}
func btnActionMoving()    {}

func (btn *Button) pressed() bool {
	btnActionPressed()
	return true
}

func (btn *Button) holding() bool {
	btnActionHolding()
	return true
}

func (btn *Button) unpressed() bool {
	btnActionUnPressed()
	return true
}

func (btn *Button) moving() bool {
	btnActionMoving()
	return true
}
