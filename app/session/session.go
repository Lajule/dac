package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lajule/dac/ent"
	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"
)

type Session struct {
	text []rune

	inputs []bool

	screen tcell.Screen

	ticker *time.Ticker

	done chan bool
}

func NewSession(text string) (*Session, error) {
	sc, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed creating screen: %w", err)
	}

	if err := sc.Init(); err != nil {
		return nil, fmt.Errorf("failed initializing screen: %w", err)
	}

	return &Session{
		text:   []rune(text),
		screen: sc,
		ticker: time.NewTicker(time.Second),
		done:   make(chan bool),
	}, nil
}

func (s *Session) Start(mu *ent.TrainingMutation) {
	go func() {
		for {
			select {
			case <-s.done:
				return

			case <-s.ticker.C:
				ev := &tcell.EventTime{}
				ev.SetEventNow()
				_ = s.screen.PostEvent(ev)
			}
		}
	}()

	for {
		switch ev := s.screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				s.inputs = append(s.inputs, s.text[len(s.inputs)] == ev.Rune())
				s.setAccuracy(mu)
				s.setProgress(mu)
				s.draw(mu)

				if len(s.inputs) == len(s.text) {
					s.stop()
					return
				}

			case tcell.KeyBackspace2:
				if len(s.inputs) > 0 {
					s.inputs = s.inputs[:len(s.inputs)-1]
					s.setAccuracy(mu)
					s.setProgress(mu)
					s.draw(mu)
				}

			case tcell.KeyEscape:
				s.stop()
				return

			case tcell.KeyCtrlL:
				s.screen.Sync()
				s.draw(mu)
			}

		case *tcell.EventTime:
			if s.isFinish(mu) {
				s.stop()
				return
			}

			s.setStopwatch(mu)
			s.setSpeed(mu)
			s.draw(mu)

		case *tcell.EventResize:
			s.screen.Sync()
			s.draw(mu)
		}
	}
}

func (s *Session) isFinish(mu *ent.TrainingMutation) bool {
	closable, _ := mu.Closable()
	duration, _ := mu.Duration()
	stopwatch, _ := mu.Stopwatch()
	return closable && duration-stopwatch == 0.0
}

func (s *Session) stop() {
	s.ticker.Stop()
	s.done <- true
	s.screen.Fini()
}

func (s *Session) setAccuracy(mu *ent.TrainingMutation) {
	var accuracy float64

	count := countValue(s.inputs, true)
	if count > 0 {
		accuracy = float64((count * 100) / len(s.inputs))
	}

	mu.SetAccuracy(accuracy)
}

func (s *Session) setProgress(mu *ent.TrainingMutation) {
	mu.SetProgress(float64((len(s.inputs) * 100) / len(s.text)))
}

func (s *Session) setStopwatch(mu *ent.TrainingMutation) {
	stopwatch, _ := mu.Stopwatch()
	stopwatch += 1
	mu.SetStopwatch(stopwatch)
}

func (s *Session) setSpeed(mu *ent.TrainingMutation) {
	if len(s.inputs) > 0 {
		index := len(s.inputs)
		if index < len(s.text)-1 && s.text[index+1] != ' ' {
			for {
				if index == 0 || s.text[index] == ' ' {
					break
				}
				index -= 1
			}
		}

		words := strings.Fields(string(s.text[:index]))
		if len(words) > 0 {
			stopwatch, _ := mu.Stopwatch()
			mu.SetSpeed(float64(len(words)*60) / stopwatch)
		}
	} else {
		mu.SetSpeed(0.0)
	}
}

func (s *Session) draw(mu *ent.TrainingMutation) {
	s.screen.Clear()
	s.drawStatus(mu)
	s.drawText()
	s.screen.Show()
}

func (s *Session) drawStatus(mu *ent.TrainingMutation) {
	st := &Status{
		screen: s.screen,
	}
	st.Draw(mu)
}

func (s *Session) drawText() {
	tx := &Text{
		screen: s.screen,
		y:      1,
	}
	tx.w, tx.h = s.screen.Size()
	tx.max = tx.w * (tx.h - tx.y)
	tx.Draw(s.text, s.inputs)
}

func countValue[T comparable](items []T, val T) (count int) {
	for _, item := range items {
		if item == val {
			count += 1
		}
	}
	return
}
