package main

import (
	"context"
	"embed"
	"log"
	"scoreboard-manager/events"
	"scoreboard-manager/scoreboard"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	scoreBoard := scoreboard.NewScoreboard(3, "round1", []string{"John", "Paul"})

	err := wails.Run(&options.App{
		Title:    "wails-test",
		Width:    900,
		Height:   600,
		LogLevel: logger.DEBUG,
		OnStartup: func(ctx context.Context) {
			scoreBoard.Startup(ctx)
		},
		WindowStartState: options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			scoreBoard,
		},
		EnumBind: []interface{}{
			events.Events,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Icon: icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
