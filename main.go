package main

import (
	"context"
	"embed"
	"log"
	"os"

	"github.com/rajatxs/go-fconsole/config"
	"github.com/rajatxs/go-fconsole/services"
	"github.com/rajatxs/go-fconsole/util"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func preconfig() error {
	var err error

	// parse default config
	if err = config.Parse(); err != nil {
		log.Fatal(err)
	}

	// create the root configuration directory if it does not exist
	configDir := config.RootDir()
	_, err = os.Stat(configDir)

	if os.IsNotExist(err) {
		if err = os.Mkdir(configDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	util.InitLogger()
	return err
}

func runApp() error {
	// Create an instance of the app structure
	app := NewApp()

	// Create service instances
	postService := services.NewPostService()
	topicService := services.NewTopicService()

	// Create application with options
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
			postService.TopicServiceRef = topicService
		},
		OnShutdown: app.terminate,
		Bind: []interface{}{
			app,
			postService,
			topicService,
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
	util.Attempt(preconfig())
	util.Attempt(runApp())
}
