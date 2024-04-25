package SysOps

import (
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func initLinux() *linux.Options {

	return &linux.Options{
		Icon:                nil,
		WindowIsTranslucent: false,
		Messages:            nil,
		WebviewGpuPolicy:    0,
		ProgramName:         "",
	}
}
func initWindows() *windows.Options {

	return &windows.Options{
		WebviewIsTransparent:                false,
		WindowIsTranslucent:                 false,
		DisableWindowIcon:                   false,
		IsZoomControlEnabled:                false,
		ZoomFactor:                          0,
		DisablePinchZoom:                    false,
		DisableFramelessWindowDecorations:   false,
		WebviewUserDataPath:                 "",
		WebviewBrowserPath:                  "",
		Theme:                               0,
		CustomTheme:                         nil,
		BackdropType:                        0,
		Messages:                            nil,
		ResizeDebounceMS:                    0,
		OnSuspend:                           nil,
		OnResume:                            nil,
		WebviewGpuIsDisabled:                false,
		WebviewDisableRendererCodeIntegrity: false,
		EnableSwipeGestures:                 false,
	}
}
func initMaosx() *mac.Options {

	return &mac.Options{
		TitleBar:             nil,
		Appearance:           "",
		WebviewIsTransparent: false,
		WindowIsTranslucent:  false,
		Preferences:          nil,
		DisableZoom:          false,
		About:                nil,
		OnFileOpen:           nil,
		OnUrlOpen:            nil,
	}
}
