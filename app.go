package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"scoreboard-manager/scoreboard"
	"strings"
	"time"
)

type App struct {
	ctx context.Context
	str string
}

func NewApp() *App {
	return &App{
		str: "lala",
	}
}

func (a *App) startup(ctx context.Context, s *scoreboard.Scoreboard) {
	a.ctx = ctx
	s.SaveEntireScoreboard()
}

func (a *App) domReady(ctx context.Context) {}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) shutdown(ctx context.Context) {}

func (a *App) GetStr() string {
	return a.str
}

func (a *App) ChangeStr() {
	fmt.Println("change str")
	if a.str == "lala" {
		a.str = "lolo"
	} else {
		a.str = "lala"
	}
	time.Sleep(1 * time.Second)
	fmt.Println("emit event")
	runtime.EventsEmit(a.ctx, string(StringChanged), a.str)
}

func (a *App) WriteStringToFile(filename string, s string) error {
	path := strings.Split(filename, "/")
	if len(path) > 1 {
		dir := strings.Join(path[:len(path)-1], "/")
		if err := os.MkdirAll(dir, 0666); err != nil {
			return err
		}
	}
	err := os.WriteFile(filename, []byte(s), 0666)
	fmt.Println("write file", filename)
	return err
}
