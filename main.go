package main

import (
	fileService "RaceSync/services"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	fileService := fileService.New()

	err := wails.Run(&options.App{
		Title:  "RaceSync",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 2, G: 6, B: 23, A: 1},
		OnDomReady: func(ctx context.Context) {
			app.startup(ctx)
			fileService.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			fileService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
