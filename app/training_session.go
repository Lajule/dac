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

func (ts *TrainingSession) Start(mu *ent.TrainingMutation) {
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
				ts.setAccuracy(mu)
				ts.setProgress(mu)
				ts.draw(mu)

				if len(ts.inputs) == len(ts.text) {
					ts.stop()
					return
				}

			case tcell.KeyBackspace2:
				if len(ts.inputs) > 0 {
					ts.inputs = ts.inputs[:len(ts.inputs)-1]
					ts.setAccuracy(mu)
					ts.setProgress(mu)
					ts.draw(mu)
				}

			case tcell.KeyEscape:
				ts.stop()
				return

			case tcell.KeyCtrlL:
				ts.screen.Sync()
				ts.draw(mu)
			}

		case *tcell.EventTime:
			if ts.isFinish(mu) {
				ts.stop()
				return
			}

			ts.setStopwatch(mu)
			ts.setSpeed(mu)
			ts.draw(mu)

		case *tcell.EventResize:
			ts.screen.Sync()
			ts.draw(mu)
		}
	}
}

func (ts *TrainingSession) isFinish(mu *ent.TrainingMutation) bool {
	closable, _ := mu.Closable()
	duration, _ := mu.Duration()
	stopwatch, _ := mu.Stopwatch()
	return closable && duration-stopwatch == 0.0
}

func (ts *TrainingSession) stop() {
	ts.ticker.Stop()
	ts.done <- true
	ts.screen.Fini()
}

func (ts *TrainingSession) setAccuracy(mu *ent.TrainingMutation) {
	var accuracy float64

	count := countValue(ts.inputs, true)
	if count > 0 {
		accuracy = float64((count * 100) / len(ts.inputs))
	}

	mu.SetAccuracy(accuracy)
}

func (ts *TrainingSession) setProgress(mu *ent.TrainingMutation) {
	mu.SetProgress(float64((len(ts.inputs) * 100) / len(ts.text)))
}

func (ts *TrainingSession) setStopwatch(mu *ent.TrainingMutation) {
	stopwatch, _ := mu.Stopwatch()
	stopwatch += 1
	mu.SetStopwatch(stopwatch)
}

func (ts *TrainingSession) setSpeed(mu *ent.TrainingMutation) {
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
			stopwatch, _ := mu.Stopwatch()
			mu.SetSpeed(float64(len(words)*60) / stopwatch)
		}
	} else {
		mu.SetSpeed(0.0)
	}
}

func (ts *TrainingSession) draw(mu *ent.TrainingMutation) {
	ts.screen.Clear()
	ts.drawStatus(mu)
	ts.drawText()
	ts.screen.Show()
}

func (ts *TrainingSession) drawStatus(mu *ent.TrainingMutation) {
	st := &Status{
		screen: ts.screen,
	}
	st.Draw(mu)
}

func (ts *TrainingSession) drawText() {
	tx := &Text{
		screen: ts.screen,
		y:      1,
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
