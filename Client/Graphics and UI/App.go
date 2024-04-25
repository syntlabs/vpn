package SysOps

import (
	"context"
	"fmt"
	wails2 "github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"time"
)

type App struct {
	ctx context.Context
}

type fnApp interface {
	firstLoad(context.Context)
	infinityPolling()
	startup(context.Context)
	shutdown(context.Context)
}

func (b *App) firstLoad(ctx context.Context) {
	b.ctx = ctx
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx

	b.infinitiPolling()
}

func (b *App) infinitiPolling() {
	ch := make(chan bool)

	go func() {
		for {
			ch <- true
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			<-ch
			fmt.Println("App is running")
			time.Sleep(1 * time.Second)
		}
	}()

	select {}
}

func (b *App) shutdown(ctx context.Context) {

}

func main() {

	allDims := UISet{}
	allDims.init()

	appBackgroundColor := options.NewRGBA(
		allDims.v["APP"]["START_BACKGROUND_COLOR"].([]uint8)[0],
		allDims.v["APP"]["START_BACKGROUND_COLOR"].([]uint8)[1],
		allDims.v["APP"]["START_BACKGROUND_COLOR"].([]uint8)[2],
		allDims.v["APP"]["START_BACKGROUND_COLOR"].([]uint8)[3])

	_ = wails2.Run(&options.App{
		Title:            "Vpn client app",
		Width:            allDims.v["APP"]["SCREEN_WIDTH"].(int),
		Height:           allDims.v["APP"]["SCREEN_HEIGHT"].(int),
		MinWidth:         allDims.v["APP"]["MIN_WIDTH"].(int),
		MinHeight:        allDims.v["APP"]["MIN_HEIGHT"].(int),
		MaxWidth:         allDims.v["APP"]["MAX_WIDTH"].(int),
		MaxHeight:        allDims.v["APP"]["MAX_HEIGHT"].(int),
		BackgroundColour: appBackgroundColor,
		AssetServer:      nil,
		Menu:             nil,
		Windows:          initWindows(),
		Mac:              initMaosx(),
		Linux:            initLinux(),
	})
}
