package SysOps

import (
	"github.com/go-playground/colors"
)

type Object struct {
	uid string
}

type Button struct {
	Object
	onFrame string
	text    string
	dims    map[string]string
	color   colors.HEXColor
}

func (btn *Button) pressed() {

	switch btn.uid {
	case "cancel":
		//
	case "startConn":
		//
	}
}

func (btn *Button) moving() {
	switch btn.uid {
	case "cancel":
		//
	case "startConn":
		//
	}
}
