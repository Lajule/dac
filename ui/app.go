package ui

import (
	"fmt"
	"sync"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type App struct {
	Duration time.Duration

	Text []rune

	Input []bool

	screen tcell.Screen

	ticker *time.Ticker

	done chan bool

	mu sync.RWMutex
}

func NewApp(duration time.Duration, text string) (*App, error) {
	encoding.Register()
	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}

	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	return &App{
		Duration: duration,
		Text:     []rune(text),
		screen:   sc,
		ticker:   time.NewTicker(time.Second),
		done:     make(chan bool),
	}, nil
}

func (app *App) Start() {
	go func() {
		for {
			select {
			case <-app.done:
				return
			case <-app.ticker.C:
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
				if app.last() {
					app.ticker.Stop()
					app.done <- true
					app.screen.Fini()
					return
				}
				app.append(ev.Rune())
				app.draw()
			case tcell.KeyEscape:
				app.ticker.Stop()
				app.done <- true
				app.screen.Fini()
				return
			case tcell.KeyCtrlL:
				app.screen.Sync()
				app.draw()
			}
		case *tcell.EventResize:
			app.screen.Sync()
			app.draw()
		}
	}
}

func (app *App) tic() {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.Duration -= time.Second
}

func (app *App) append(r rune) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.Input = append(app.Input, app.Text[len(app.Input)] == r)
}

func (app *App) last() bool {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return len(app.Input) == len(app.Text)
}

func (app *App) draw() {
	app.mu.RLock()
	defer app.mu.RUnlock()

	w, h := app.screen.Size()
	max := w * (h - 1)

	app.screen.Clear()

	if app.Duration > 0 {
		duration := app.Duration.String()
		for i, r := range []rune(duration) {
			x := w - len(duration) + i
			app.screen.SetContent(x, 0, r, nil, tcell.StyleDefault)
		}
	}

	textChunks := chunkBy(app.Text, max)
	inputChunks := chunkBy(app.Input, max)
	offset := len(app.Input) / max

	y := 1
	for i, r := range textChunks[offset] {
		x := i % w
		if i > 0 && x == 0 {
			if y == h {
				break
			}
			y += 1
		}
		style := tcell.StyleDefault
		if len(inputChunks) > offset && i < len(inputChunks[offset]) {
			style = style.Bold(true)
			if inputChunks[offset][i] {
				style = style.Foreground(tcell.ColorGreen)
			} else {
				style = style.Foreground(tcell.ColorRed)
			}
			if unicode.IsSpace(r) {
				style = style.Underline(true)
			}
		}
		app.screen.SetContent(x, y, r, nil, style)
	}

	app.screen.Show()
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
