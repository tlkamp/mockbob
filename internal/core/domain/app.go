package domain

import "github.com/tlkamp/mockbob/internal/core/ports"

type App struct {
	bob ports.Bob
}

func NewApp(bob ports.Bob) *App {
	return &App{bob: bob}
}

func (a *App) Process(input string) string {
	return a.bob.Bobify(input)
}
