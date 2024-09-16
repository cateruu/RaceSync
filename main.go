package main

import (
	"embed"
	"log"
	"time"

	"github.com/cateruu/iRacing-utility/internal/fileservice"
	"github.com/cateruu/iRacing-utility/internal/overlayservice"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/plugins/experimental/single_instance"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "iRacing utility",
		Description: "iRacing utility tool",
		Services: []application.Service{
			application.NewService(fileservice.New()),
			application.NewService(overlayservice.New()),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Plugins: map[string]application.Plugin{
			"single_instance": single_instance.NewPlugin(&single_instance.Config{
				ActivateAppOnSubsequentLaunch: true,
			}),
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:            "iRacing Utility",
		BackgroundColour: application.NewRGB(0, 0, 0),
		URL:              "/",
		MinWidth:         800,
		MinHeight:        600,
	})

	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Events.Emit(&application.WailsEvent{
				Name: "time",
				Data: now,
			})
			time.Sleep(time.Second)
		}
	}()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
