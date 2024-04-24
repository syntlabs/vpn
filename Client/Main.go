package Utils

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"time"
)

const JS_PATH string = ""
const TS_PATH string = ""
const CSS_PATH string = ""
const MAIN_URL_PATH string = ""

var dIdle = 1 * time.Second.Seconds()
var dThreads uint8 = 1
var dMRetries uint8 = 10
var dUpTime time.Duration = 1

func main() {
	if firstRun() {
		show_greetings()
		Client.createMainClient()
	} else {
		proccess_standart_run()
	}

	n_daemon := NetworkDaemon{
		endpoint:   MAIN_URL_PATH,
		idleTime:   dIdle,
		updateTime: dUpTime,
		threads:    dThreads,
		maxRetries: dMRetries,
		state: "is_active"
	}

	n_daemon.run()
}

func first_ever_run() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	stats := &sys.Stats{}

	globalApp := wails.CreateApp(&wails.AppConfig{
		Width:  512,
		Height: 512,
		Title:  "CPU Usage",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	globalApp.Bind(stats)
	globalApp.Run()
}

func proccess_standart_run() {}

func show_greetings() {}
