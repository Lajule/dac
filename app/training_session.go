package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lajule/dac/ent"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type TrainingSession struct {
	text []rune

	inputs []bool

	screen tcell.Screen

	ticker *time.Ticker

	done chan bool
}

func NewTrainingSession(text string) (*TrainingSession, error) {
	encoding.Register()

	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}

	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	return &TrainingSession{
		text:   []rune(text),
		screen: sc,
		ticker: time.NewTicker(time.Second),
		done:   make(chan bool),
	}, nil
}

func (ts *TrainingSession) Start(t *ent.TrainingMutation) {
	go func() {
		for {
			select {
			case <-ts.done:
				return

			case <-ts.ticker.C:
				ev := &tcell.EventTime{}
				ev.SetEventNow()
				ts.screen.PostEvent(ev)
			}
		}
	}()

	for {
		switch ev := ts.screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				ts.inputs = append(ts.inputs, ts.text[len(ts.inputs)] == ev.Rune())
				ts.setAccuracy(t)
				ts.setProgress(t)
				ts.draw(t)
				if len(ts.inputs) == len(ts.text) {
					ts.stop()
					return
				}

			case tcell.KeyBackspace2:
				if len(ts.inputs) > 0 {
					ts.inputs = ts.inputs[:len(ts.inputs)-1]
					ts.setAccuracy(t)
					ts.setProgress(t)
					ts.draw(t)
				}

			case tcell.KeyEscape:
				ts.stop()
				return

			case tcell.KeyCtrlL:
				ts.screen.Sync()
				ts.draw(t)
			}

		case *tcell.EventTime:
			closable, _ := t.Closable()
			duration, _ := t.Duration()
			stopwatch, _ := t.AddedStopwatch()
			if closable && duration-stopwatch == 0.0 {
				ts.stop()
				return
			}
			ts.setSpeed(t)
			ts.draw(t)

		case *tcell.EventResize:
			ts.screen.Sync()
			ts.draw(t)
		}
	}
}

func (ts *TrainingSession) stop() {
	ts.ticker.Stop()
	ts.done <- true
	ts.screen.Fini()
}

func (ts *TrainingSession) setAccuracy(t *ent.TrainingMutation) {
	count := countValue(ts.inputs, true)
	if count > 0 {
		t.SetAccuracy(float64((count * 100) / len(ts.inputs)))
	} else {
		t.SetAccuracy(0.0)
	}
}

func (ts *TrainingSession) setProgress(t *ent.TrainingMutation) {
	t.SetProgress(float64((len(ts.inputs) * 100) / len(ts.text)))
}

func (ts *TrainingSession) setSpeed(t *ent.TrainingMutation) {
	t.AddStopwatch(1.0)
	if len(ts.inputs) > 0 {
		index := len(ts.inputs)
		if index < len(ts.text)-1 && ts.text[index+1] != ' ' {
			for {
				if index == 0 || ts.text[index] == ' ' {
					break
				}
				index -= 1
			}
		}

		words := strings.Fields(string(ts.text[:index]))
		if len(words) > 0 {
			stopwatch, _ := t.AddedStopwatch()
			t.SetSpeed(float64(len(words)*60) / stopwatch)
		}
	} else {
		t.SetSpeed(0.0)
	}
}

func (ts *TrainingSession) draw(t *ent.TrainingMutation) {
	ts.screen.Clear()
	ts.drawStatus(t)
	ts.drawText()
	ts.screen.Show()
}

func (ts *TrainingSession) drawStatus(t *ent.TrainingMutation) {
	st := &Status{
		screen: ts.screen,
	}
	st.Draw(t)
}

func (ts *TrainingSession) drawText() {
	tx := &Text{
		screen: ts.screen,
		y: 1,
	}
	tx.w, tx.h = ts.screen.Size()
	tx.max = tx.w * (tx.h - tx.y)
	tx.Draw(ts.text, ts.inputs)
}

func countValue[T comparable](items []T, val T) (count int) {
	for _, item := range items {
		if item == val {
			count += 1
		}
	}
	return
}
