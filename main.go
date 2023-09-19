package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func runApp() error {
	// Create an instance of the app structure
	app := NewApp()

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
		OnStartup:  app.startup,
		OnShutdown: app.terminate,
		Bind: []interface{}{
			app,
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
