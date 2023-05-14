package app

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type Dac struct {
	Duration time.Duration

	Text []rune

	Inputs []bool

	Precision int

	Speed int

	screen tcell.Screen

	ticker *time.Ticker

	done chan bool
}

func NewDac(duration time.Duration, text string) (*Dac, error) {
	encoding.Register()
	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}

	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	return &Dac{
		Duration: duration,
		Text:     []rune(text),
		screen:   sc,
		ticker:   time.NewTicker(time.Second),
		done:     make(chan bool),
	}, nil
}

func (d *Dac) Start() {
	tics := 0
	go func() {
		for {
			select {
			case <-d.done:
				return
			case <-d.ticker.C:
				ev := &tcell.EventTime{}
				ev.SetEventNow()
				d.screen.PostEvent(ev)
			}
		}
	}()

	for {
		switch ev := d.screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				if len(d.Inputs) == len(d.Text) {
					d.stop()
					return
				}
				d.Inputs = append(d.Inputs, d.Text[len(d.Inputs)] == ev.Rune())
				d.precision()
				d.draw()
			case tcell.KeyBackspace2:
				if len(d.Inputs) > 0 {
					d.Inputs = d.Inputs[:len(d.Inputs)-1]
					d.precision()
					d.draw()
				}
			case tcell.KeyEscape:
				d.stop()
				return
			case tcell.KeyCtrlL:
				d.screen.Sync()
				d.draw()
			}
		case *tcell.EventTime:
			if d.Duration == 0 {
				d.stop()
				return
			}
			tics += 1
			d.Duration -= time.Second
			if len(d.Inputs) > 0 {
				index := len(d.Inputs)
				if index < len(d.Text)-1 && !unicode.IsSpace(d.Text[index+1]) {
					for {
						if index == 0 || unicode.IsSpace(d.Text[index]) {
							break
						}
						index -= 1
					}
				}
				words := strings.Fields(string(d.Text[:index]))
				if len(words) > 0 {
					d.Speed = (len(words) * 60) / tics
				}
			} else {
				d.Speed = 0
			}
			d.draw()
		case *tcell.EventResize:
			d.screen.Sync()
			d.draw()
		}
	}
}

func (d *Dac) stop() {
	d.ticker.Stop()
	d.done <- true
	d.screen.Fini()
}

func (d *Dac) precision() {
	count := countVal(d.Inputs, true)
	if count > 0 {
		d.Precision = (count * 100) / len(d.Inputs)
	} else {
		d.Precision = 0
	}
}

func (d *Dac) draw() {
	w, h := d.screen.Size()
	max := w * (h - 1)

	d.screen.Clear()

	status := fmt.Sprintf("(%d,%d) %d%% %dw/s %s", len(d.Text), len(d.Inputs), d.Precision, d.Speed, d.Duration.String())
	for i, r := range []rune(fmt.Sprintf("%*s", w, status)) {
		d.screen.SetContent(i, 0, r, nil, tcell.StyleDefault.Bold(true).Reverse(true))
	}

	textChunks := chunkBy(d.Text, max)
	inputChunks := chunkBy(d.Inputs, max)
	offset := len(d.Inputs) / max

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
		d.screen.SetContent(x, y, r, nil, style)
	}

	d.screen.Show()
}

func countVal[T comparable](items []T, val T) (count int) {
	for _, item := range items {
		if item == val {
			count += 1
		}
	}
	return
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
