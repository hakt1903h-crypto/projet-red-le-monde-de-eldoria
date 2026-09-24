package ui

import (
	"time"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewMainMenu(app *tview.Application) *tview.Pages {

	background := NewASCIIBackground(app)

	title := tview.NewTextView()
	title.SetText(`
   ███████╗██╗     ██████╗  ██████╗ ██████╗ ██████╗ ██╗ █████╗
   ██╔════╝██║     ██╔══██╗██╔═══██╗██╔══██╗██╔══██╗██║██╔══██╗
   █████╗  ██║     ██║  ██║██║   ██║██████╔╝██████╔╝██║███████║
   ██╔══╝  ██║     ██║  ██║██║   ██║██╔══██╗██╔══██╗██║██╔══██║
   ██║     ███████╗██████╔╝╚██████╔╝██║  ██║██║  ██║██║██║  ██║
   ╚═╝     ╚══════╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝
	`)

	title.SetTextAlign(tview.AlignCenter)
	title.SetTextColor(ColorPrimary)
	title.SetBackgroundColor(ColorBackground)

	menu := createMainMenu(app)

	center := tview.NewFlex()
	center.SetDirection(tview.FlexRow)

	center.AddItem(nil, 5, 0, false)
	center.AddItem(title, 9, 0, false)
	center.AddItem(nil, 2, 0, false)
	center.AddItem(menu, 9, 0, true)
	center.AddItem(nil, 5, 0, false)

	root := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(center, 54, 0, true).
		AddItem(nil, 0, 1, false)

	pages := tview.NewPages()
	pages.AddPage("background", background, true, true)
	pages.AddPage("menu", root, true, true)

	return pages
}

func createMainMenu(app *tview.Application) *tview.List {

	menu := tview.NewList()

	menu.ShowSecondaryText(false)

	menu.SetMainTextColor(ColorText)
	menu.SetSelectedTextColor(tcell.ColorWhite)
	menu.SetSelectedBackgroundColor(ColorSelected)

	menu.SetHighlightFullLine(true)
	menu.SetSelectedFocusOnly(true)

	menu.SetBorder(true)
	menu.SetBorderColor(ColorBorder)
	menu.SetTitle(" MENU ")
	menu.SetTitleColor(ColorPrimary)
	menu.SetTitleAlign(tview.AlignCenter)

	menu.AddItem(
		"▶  JOUER",
		"",
		0,
		func() {
			startGame(app)
		},
	)

	menu.AddItem(
		"⚙  OPTIONS",
		"",
		0,
		func() {
			openOptions(app)
		},
	)

	menu.AddItem(
		"✕  QUITTER",
		"",
		0,
		func() {
			app.Stop()
		},
	)

	return menu
}

func startGame(app *tview.Application) {
	// Plus tard :
	// → création personnage
	// → exploration
	// → etc.
}

func openOptions(app *tview.Application) {
	// Plus tard :
	// → volume musique
	// → retour
}

func NewASCIIBackground(app *tview.Application) *tview.TextView {

	video, err := NewASCIIVideo(
		app,
		"assets/ascii/intro",
		80*time.Millisecond,
	)

	if err != nil {

		background := tview.NewTextView()

		background.SetText(
			"\n\n\n\n\n\n\n\n\n" +
				"          ELDORIA\n\n" +
				"   Les Fragments du Souvenir",
		)

		background.SetTextAlign(tview.AlignCenter)
		background.SetBackgroundColor(ColorBackground)
		background.SetTextColor(ColorSecondary)

		return background
	}

	video.Start()

	return video.View()
}
