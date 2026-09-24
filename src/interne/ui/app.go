package ui

import "github.com/rivo/tview"

type App struct {
	Application *tview.Application
}

func NewApp() *App {
	return &App{
		Application: tview.NewApplication(),
	}
}

func (app *App) Run() error {
	root := NewMainMenu(app.Application)

	app.Application.SetRoot(root, true)

	return app.Application.Run()
}
