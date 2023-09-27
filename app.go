package main

import (
	"context"
	"log"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/rajatxs/go-fconsole/config"
	"github.com/rajatxs/go-fconsole/db"
	"github.com/rajatxs/go-fconsole/types"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	var err error

	a.ctx = ctx
	if err = db.ConnectMongoDb(ctx); err != nil {
		log.Fatal(err)
	}
}

// terminate is called when the app shutdown.
func (a *App) terminate(ctx context.Context) {
	var err error

	if err = db.DisconnectMongoDb(ctx); err != nil {
		log.Println(err)
	}
}

// OpenBrowser opens given url into default browser
func (a *App) OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}

// GetAppConfigVariables returns public app config variables
func (a *App) GetAppConfigVariables() (env *types.AppPublicConfigVariables) {
	env = &types.AppPublicConfigVariables{}

	env.ENV = config.Env()
	env.ADMIN_ID = config.AdminId()
	env.CLOUDINARY_ID = config.CloudinaryId()
	return env
}

// GetVersions returns app version information
func (a *App) GetVersions() (ver *types.AppVersions) {
	ver = &types.AppVersions{}

	currentUser, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	ver.App = "0.0.1"
	ver.Date = "2023-09-20T07:53:14.501Z"
	ver.Wails = "2.6.0"
	ver.Go = "1.20.3"
	ver.WebView2 = "111.0.2045.32"
	ver.Os = runtime.GOOS
	ver.Arch = runtime.GOARCH
	ver.Username = currentUser.Name
	ver.HomeDir = currentUser.HomeDir
	return ver
}
