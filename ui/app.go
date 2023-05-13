package ui

import (
	"fmt"
	"sync"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

var (
	inputStyles = map[bool]tcell.Style{
		true:  tcell.StyleDefault.Foreground(tcell.ColorGreen).Bold(true),
		false: tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true),
	}
)

type App struct {
	screen tcell.Screen

	duration time.Duration

	ticker *time.Ticker

	done chan bool

	text []rune

	input []bool

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
		screen:   sc,
		duration: duration,
		ticker:   time.NewTicker(time.Second),
		done:     make(chan bool),
		text:     []rune(text),
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

	app.duration -= time.Second
}

func (app *App) append(r rune) {
	app.mu.Lock()
	defer app.mu.Unlock()

	app.input = append(app.input, app.text[len(app.input)] == r)
}

func (app *App) last() bool {
	app.mu.RLock()
	defer app.mu.RUnlock()

	return len(app.input) == len(app.text)
}

func (app *App) draw() {
	app.mu.RLock()
	defer app.mu.RUnlock()

	w, h := app.screen.Size()
	max := w * (h - 1)

	app.screen.Clear()

	duration := app.duration.String()
	for i, r := range []rune(duration) {
		x := w - len(duration) + i

		app.screen.SetContent(x, 0, r, nil, tcell.StyleDefault)
	}

	textChunks := chunkBy(app.text, max)
	inputChunks := chunkBy(app.input, max)
	offset := len(app.input) / max

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
			style = inputStyles[inputChunks[offset][i]]

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
