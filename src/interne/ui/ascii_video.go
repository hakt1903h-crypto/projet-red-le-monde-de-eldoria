package ui

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/rivo/tview"
)

type ASCIIVideo struct {
	frames   []string
	view     *tview.TextView
	app      *tview.Application
	interval time.Duration
}

func NewASCIIVideo(
	app *tview.Application,
	path string,
	interval time.Duration,
) (*ASCIIVideo, error) {

	files, err := filepath.Glob(filepath.Join(path, "*.txt"))
	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	frames := make([]string, 0, len(files))

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		frames = append(
			frames,
			strings.TrimSuffix(string(data), "\n"),
		)
	}

	view := tview.NewTextView()

	view.SetTextAlign(tview.AlignCenter)
	view.SetScrollable(false)
	view.SetWrap(false)

	view.SetBackgroundColor(ColorBackground)
	view.SetTextColor(ColorSecondary)

	return &ASCIIVideo{
		frames:   frames,
		view:     view,
		app:      app,
		interval: interval,
	}, nil
}

func (video *ASCIIVideo) View() *tview.TextView {
	return video.view
}

func (video *ASCIIVideo) Start() {

	if len(video.frames) == 0 {
		return
	}

	go func() {

		ticker := time.NewTicker(video.interval)
		defer ticker.Stop()

		index := 0

		for range ticker.C {

			frame := video.frames[index]

			video.app.QueueUpdateDraw(func() {
				video.view.SetText(frame)
			})

			index++

			if index >= len(video.frames) {
				index = 0
			}
		}

	}()
}
