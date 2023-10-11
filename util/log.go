package util

import (
	"path/filepath"

	"github.com/rajatxs/go-fconsole/config"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

var Log logger.Logger

func InitLogger() {
	file := filepath.Join(config.RootDir(), "debug.log")
	Log = logger.NewFileLogger(file)
}
