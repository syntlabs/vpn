package main

import "github.com/go-playground/colors"

type EventListener interface {
	pressed() bool
	holding() bool
	unpressed() bool
	moving() bool
}

type Button struct {
	Object
	onFrame string
	text  string
	dims  map[string]string
	color colors.HEXColor
	event EventListener
}


func buttonPressed(btn *Button) {

	var testButtonIds []string = []string {
		"Main_01", "Main_02", "Detail_01"
	}

	switch btn.obid {
	case "Main_01":
		newlayout = ///
		testBackground := fglob(id: "Background_01")
		newColor := Color.marine
		changeLayout(testBackground.color, newColor)
	case "Main_02":
		data := map[string]interface{
			"somedata": 14213,
			"somedata2":  "Hello World!",
		}
		sendData(data: data)
	case "Main_03":
		localFrame := fglob(id: btn.onFrame)
		nextFrame := fglob(id: __Frame__.detail)
		changeLayout(from: localFrame, to: nextFrame)
}


func buttonUnpressed() {}
func buttonHolding()   {}
func buttonMoving()    {}

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

func changeLayout(referenceProperty Object, newProperty Object) {
	//frontCallback(some_event, some_data)
}

func sendData(data map[string]interface{}){
	//frontCallback(some_event, some_data)
}

func moveFrame(from: __Frame__, to __Frame__) {
	//frontCallback(some_event, some_data)
}