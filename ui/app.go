package ui

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

var (
	green = tcell.StyleDefault.Foreground(tcell.ColorGreen).Bold(true)

	red = tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true)
)

type App struct {
	screen tcell.Screen

	duration time.Duration

	text []rune

	styles []tcell.Style

	index int

	mu sync.RWMutex
}

func NewApp(text string) (*App, error) {
	encoding.Register()

	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}

	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	return &App{
		screen: sc,
		text:   []rune(text),
	}, nil
}

func (app *App) Start() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				app.tic()
				app.draw()
			}
		}
	}()

	for {
		switch ev := app.screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				app.update(ev.Rune())

			case tcell.KeyEscape:
				app.screen.Fini()
				os.Exit(0)

			case tcell.KeyCtrlL:
				app.screen.Sync()
			}

			app.draw()

		case *tcell.EventResize:
			app.screen.Sync()
			app.draw()
		}
	}
}

func (app *App) tic() {
	app.mu.Lock()
	defer app.mu.Unlock()

	app.duration += time.Second
}

func (app *App) update(r rune) {
	app.mu.Lock()
	defer app.mu.Unlock()

	app.index += 1

	if app.text[app.index-1] == r {
		app.styles = append(app.styles, green)
	} else {
		app.screen.Beep()
		app.styles = append(app.styles, red)
	}
}

func (app *App) draw() {
	app.mu.RLock()
	defer app.mu.RUnlock()

	w, h := app.screen.Size()

	app.screen.Clear()

	duration := app.duration.String()

	startAt := w - len(duration)
	for i, r := range []rune(duration) {
		x := startAt + i

		app.screen.SetContent(x, 0, r, nil, tcell.StyleDefault)
	}

	y := 1
	for i, r := range app.text {
		x := i % w

		if i > 0 && x == 0 {
			if y == h {
				break
			}

			y += 1
		}

		style := tcell.StyleDefault

		if i < len(app.styles) {
			style = app.styles[i]
		}

		app.screen.SetContent(x, y, r, nil, style)
	}

	app.screen.Show()
}
