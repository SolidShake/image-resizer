package main

import (
	"embed"
	"fmt"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/zap"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure

	// fmt.Println(runtime.OpenDialogOptions.DefaultDirectory)
	dir, err := homedir.Dir()
	fmt.Println(dir, err)

	if runtime.GOOS == "darwin" {
		dir += "/Desktop"
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	app := NewApp(sugar, dir)

	// defaultPath := "/"

	// platform := runtime.GOOS
	// switch platform {
	// case "darwin":
	// 	defaultPath := fmt.Sprintf("/Users/%s/Desktop", user)
	// }
	// fmt.Println(platform)
	// fmt.Println(defaultPath)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "image-resizer",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
