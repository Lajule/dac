package app

import (
	"fmt"
	"time"

	"github.com/Lajule/dac/ent"
	"github.com/gdamore/tcell/v2"
)

type Status struct {
	screen tcell.Screen

	x int
}

func (s *Status) Draw(t *ent.TrainingMutation) {
	s.drawAccuracy(t)
	s.drawSeparator()
	s.drawSpeed(t)
	s.drawSeparator()
	s.drawProgress(t)
	s.drawSeparator()
	s.drawStopwatch(t)
}

func (s *Status) drawSeparator() {
	s.screen.SetContent(s.x, 0, '|', nil, tcell.StyleDefault)
	s.x += 1
}

func (s *Status) drawAccuracy(t *ent.TrainingMutation) {
	accuracy, _ := t.Accuracy()
	style := tcell.StyleDefault
	if accuracy < 50.0 {
		style = style.Foreground(tcell.ColorRed)
	} else {
		style = style.Foreground(tcell.ColorGreen)
	}
	for _, r := range fmt.Sprintf("%*.0f%%", 3, accuracy) {
		s.screen.SetContent(s.x, 0, r, nil, style)
		s.x += 1
	}
}

func (s *Status) drawSpeed(t *ent.TrainingMutation) {
	speed, _ := t.Speed()
	for _, r := range fmt.Sprintf("%*.0fw/s", 3, speed) {
		s.screen.SetContent(s.x, 0, r, nil, tcell.StyleDefault)
		s.x += 1
	}
}

func (s *Status) drawProgress(t *ent.TrainingMutation) {
	progress, _ := t.Progress()
	for i := 0; i < 10; i++ {
		style := tcell.StyleDefault
		if progress/10.0 > float64(i) {
			style = style.Reverse(true)
		}
		s.screen.SetContent(s.x, 0, ' ', nil, style)
		s.x += 1
	}
}

func (s *Status) drawStopwatch(t *ent.TrainingMutation) {
	duration, _ := t.Duration()
	stopwatch, _ := t.AddedStopwatch()
	for _, r := range (time.Duration(duration-stopwatch) * time.Second).String() {
		s.screen.SetContent(s.x, 0, r, nil, tcell.StyleDefault)
		s.x += 1
	}
}
