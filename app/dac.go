package app

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/Lajule/dac/ent"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type Dac struct {
	Training *ent.Training

	text []rune

	inputs []bool

	screen tcell.Screen

	ticker *time.Ticker

	done chan bool
}

func NewDac(text string) (*Dac, error) {
	encoding.Register()
	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}
	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	return &Dac{
		text:   []rune(text),
		screen: sc,
		ticker: time.NewTicker(time.Second),
		done:   make(chan bool),
	}, nil
}

func (d *Dac) Start(t *ent.TrainingMutation) {
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
				if len(d.inputs) == len(d.text) {
					d.stop()
					return
				}
				d.inputs = append(d.inputs, d.text[len(d.inputs)] == ev.Rune())
				d.accuracy(t)
				d.progress(t)
				d.draw(t)
			case tcell.KeyBackspace2:
				if len(d.inputs) > 0 {
					d.inputs = d.inputs[:len(d.inputs)-1]
					d.accuracy(t)
					d.progress(t)
					d.draw(t)
				}
			case tcell.KeyEscape:
				d.stop()
				return
			case tcell.KeyCtrlL:
				d.screen.Sync()
				d.draw(t)
			}
		case *tcell.EventTime:
			closable, _ := t.Closable()
			duration, _ := t.Duration()
			if closable && duration == 0 {
				d.stop()
				return
			}
			t.AddStopwatch(1)
			if len(d.inputs) > 0 {
				index := len(d.inputs)
				if index < len(d.text)-1 && !unicode.IsSpace(d.text[index+1]) {
					for {
						if index == 0 || unicode.IsSpace(d.text[index]) {
							break
						}
						index -= 1
					}
				}
				words := strings.Fields(string(d.text[:index]))
				if len(words) > 0 {
					stopwatch, _ := t.AddedStopwatch()
					t.SetSpeed((len(words) * 60) / stopwatch)
				}
			} else {
				t.SetSpeed(0)
			}
			d.draw(t)
		case *tcell.EventResize:
			d.screen.Sync()
			d.draw(t)
		}
	}
}

func (d *Dac) stop() {
	d.ticker.Stop()
	d.done <- true
	d.screen.Fini()
}

func (d *Dac) progress(t *ent.TrainingMutation) {
	t.SetProgress((len(d.inputs) * 100) / len(d.text))
}

func (d *Dac) accuracy(t *ent.TrainingMutation) {
	count := countValue(d.inputs, true)
	if count > 0 {
		t.SetAccuracy((count * 100) / len(d.inputs))
	} else {
		t.SetAccuracy(0)
	}
}

func (d *Dac) status(t *ent.TrainingMutation) string {
	duration, _ := t.Duration()
	stopwatch, _ := t.Stopwatch()
	accuracy, _ := t.Accuracy()
	speed, _ := t.Speed()
	return fmt.Sprintf("(%d,%d) %d%% %dw/s %s", len(d.text), len(d.inputs), accuracy, speed, (time.Duration(duration-stopwatch) * time.Second).String())
}

func (d *Dac) draw(t *ent.TrainingMutation) {
	w, h := d.screen.Size()
	max := w * (h - 1)
	textChunks := chunkBy(d.text, max)
	inputChunks := chunkBy(d.inputs, max)
	offset := len(d.inputs) / max

	d.screen.Clear()

	style := tcell.StyleDefault.Bold(true).Reverse(true)
	for i, r := range []rune(fmt.Sprintf("%*s", w, d.status(t))) {
		d.screen.SetContent(i, 0, r, nil, style)
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

func countValue[T comparable](items []T, val T) (count int) {
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
