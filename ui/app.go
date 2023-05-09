package ui

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type App struct {
	screen tcell.Screen

	start time.Time

	input [][]rune
}

func NewApp(input io.Reader) (*App, error) {
	encoding.Register()

	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}

	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	b, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("failed reading input: %w", err)
	}

	fields := strings.Fields(string(b))

	app := &App{
		screen: sc,
		start:  time.Now(),
		input: make([][]rune, len(fields)),
	}

	for i, field := range fields {
		app.input[i] = []rune(field)
	}

	return app, nil
}

func (app *App) Start() {
	for {
		switch ev := app.screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
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

func (app *App) draw() {
	//w, h := app.screen.Size()
	style := tcell.StyleDefault
	//green := style.Foreground(tcell.ColorGreen).Bold(true)

	app.screen.Clear()

	l := 1
	for _, word := range app.input {
		for _, character := range word {
			app.screen.SetContent(l, 1, character, nil, style)
			l += 1
		}
	}

	app.screen.Show()
}
