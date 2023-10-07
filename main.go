package main

import (
	"context"
	"embed"

	"github.com/rajatxs/go-fconsole/services"
	"github.com/rajatxs/go-fconsole/util"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func runApp() error {
	// Create an instance of the app structure
	app := NewApp()

	// Create service instances
	postService := services.NewPostService()

	// Create application with options
	// TODO: Add `backgroundColor` property
	err := wails.Run(&options.App{
		Title:         "Console",
		Width:         1300,
		Height:        800,
		MinWidth:      700,
		MinHeight:     500,
		DisableResize: false,
		Fullscreen:    false,
		Frameless:     false,
		StartHidden:   false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			postService.Ctx = ctx
		},
		OnShutdown: app.terminate,
		Bind: []interface{}{
			app,
			postService,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    true,
		},
	})

	return err
}

func main() {
	util.Attempt(runApp())
}
