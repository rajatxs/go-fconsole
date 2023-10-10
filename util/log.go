package util

import "github.com/wailsapp/wails/v2/pkg/logger"

var Log logger.Logger

func init() {
	Log = logger.NewDefaultLogger()
}
