package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"ibis/backend/db"
	"ibis/backend"
)

//go:embed all:frontend/dist
var assets embed.FS

var appName = "Ibis"

func main() {
	var err error
	app := application.NewApp()

	db.InitializeDB(appName)
	
	err = wails.Run(&options.App{
		Title:  "Ibis",
		Width:  1024,
		Height: 1024,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
