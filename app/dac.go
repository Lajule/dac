package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lajule/dac/ent"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type Dac struct {
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
				d.inputs = append(d.inputs, d.text[len(d.inputs)] == ev.Rune())
				d.setAccuracy(t)
				d.setProgress(t)
				d.draw(t)
				if len(d.inputs) == len(d.text) {
					d.stop()
					return
				}
			case tcell.KeyBackspace2:
				if len(d.inputs) > 0 {
					d.inputs = d.inputs[:len(d.inputs)-1]
					d.setAccuracy(t)
					d.setProgress(t)
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
			stopwatch, _ := t.AddedStopwatch()
			if closable && duration-stopwatch == 0.0 {
				d.stop()
				return
			}
			d.setSpeed(t)
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

func (d *Dac) setAccuracy(t *ent.TrainingMutation) {
	count := countValue(d.inputs, true)
	if count > 0 {
		t.SetAccuracy(float64((count * 100) / len(d.inputs)))
	} else {
		t.SetAccuracy(0.0)
	}
}

func (d *Dac) setProgress(t *ent.TrainingMutation) {
	t.SetProgress(float64((len(d.inputs) * 100) / len(d.text)))
}

func (d *Dac) setSpeed(t *ent.TrainingMutation) {
	t.AddStopwatch(1.0)
	if len(d.inputs) > 0 {
		index := len(d.inputs)
		if index < len(d.text)-1 && d.text[index+1] != ' ' {
			for {
				if index == 0 || d.text[index] == ' ' {
					break
				}
				index -= 1
			}
		}
		words := strings.Fields(string(d.text[:index]))
		if len(words) > 0 {
			stopwatch, _ := t.AddedStopwatch()
			t.SetSpeed(float64(len(words)*60) / stopwatch)
		}
	} else {
		t.SetSpeed(0.0)
	}
}

func (d *Dac) draw(t *ent.TrainingMutation) {
	d.screen.Clear()
	d.drawStatus(t)
	d.drawText()
	d.screen.Show()
}

func (d *Dac) drawStatus(t *ent.TrainingMutation) {
	s := &Status{
		screen: d.screen,
	}
	s.Draw(t)
}

func (d *Dac) drawText() {
	t := &Text{
		screen: d.screen,
		y: 1,
	}
	t.w, t.h = d.screen.Size()
	t.max = t.w * (t.h - t.y)
	t.Draw(d.text, d.inputs)
}

func countValue[T comparable](items []T, val T) (count int) {
	for _, item := range items {
		if item == val {
			count += 1
		}
	}
	return
}
