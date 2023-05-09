package ui

import (
	"io"
	"fmt"
	"os"
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

	return &App{
		screen: sc,
		start:  time.Now(),
	}, nil
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
	//w, h := d.screen.Size()
	//style := tcell.StyleDefault
	//green := style.Foreground(tcell.ColorGreen).Bold(true)

	app.screen.Clear()

	app.screen.Show()
}
