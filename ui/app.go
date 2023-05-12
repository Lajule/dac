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
	inputStyles = map[bool]tcell.Style{
		true:  tcell.StyleDefault.Foreground(tcell.ColorGreen).Bold(true),
		false: tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true),
	}
)

type App struct {
	screen tcell.Screen

	duration time.Duration

	text []rune

	input []bool

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
	app.input = append(app.input, app.text[app.index-1] == r)
}

func (app *App) draw() {
	app.mu.RLock()
	defer app.mu.RUnlock()

	w, h := app.screen.Size()

	app.screen.Clear()

	duration := app.duration.String()
	durationLen := len(duration)
	for i, r := range []rune(duration) {
		x := w - durationLen + i

		app.screen.SetContent(x, 0, r, nil, tcell.StyleDefault)
	}

	max := w * (h - 1)

	textChunks := chunkBy(app.text, max)
	inputChunks := chunkBy(app.input, max)

	inputLen := len(app.input)
	offset := inputLen / max

	var inputChunk []bool
	var inputChunkLen int
	if len(inputChunks) == offset {
		inputChunkLen = 0
	} else {
		inputChunk = inputChunks[offset]
		inputChunkLen = len(inputChunk)
	}

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

		if i < inputChunkLen {
			style = inputStyles[inputChunk[i]]
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
