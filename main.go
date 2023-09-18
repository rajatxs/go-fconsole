package main

import (
	"embed"
	"log"

	"github.com/rajatxs/go-fconsole/db"
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
		MinWidth:  1000,
		Width:     1350,
		MinHeight: 600,
		Height:    900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	return err
}

func main() {
	var err error

	if err = db.ConnectMongoDb(); err != nil {
		log.Fatal(err)
	}

	if err = runApp(); err != nil {
		log.Fatal(err)
	}
}
