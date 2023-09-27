package main

import (
	"context"
	"embed"
	"log"

	"github.com/rajatxs/go-fconsole/services"
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
			var err error

			app.startup(ctx)
			postService.Ctx = ctx

			if err = services.InitCloudinary(); err != nil {
				log.Println(err)
			}
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
	var err error

	if err = runApp(); err != nil {
		log.Fatal(err)
	}
}
