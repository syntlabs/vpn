package main

type View struct {
	Object
}

type viewMethods interface {
	show(dims map[string]interface{}, msg string)
}

func (v View) show() {
	//frontCallback(some_event, some_data)
}

func viewNoInternet() {
	// function to show view to the user that there is no internet connection
}

func viewSubJustEnded() {
	// function to show view to the user that his paid subscription just ended recently
}