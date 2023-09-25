package main

import (
	"context"
	"embed"
	"log"

	"github.com/rajatxs/go-fconsole/services"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
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
		Title:     "Console",
		MinWidth:  700,
		Width:     1300,
		MinHeight: 500,
		Height:    700,
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
	})

	return err
}

func main() {
	var err error

	if err = runApp(); err != nil {
		log.Fatal(err)
	}
}
