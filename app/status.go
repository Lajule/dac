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

func (s *Status) Draw(mu *ent.TrainingMutation) {
	s.drawAccuracy(mu)
	s.drawSeparator()
	s.drawSpeed(mu)
	s.drawSeparator()
	s.drawProgress(mu)
	s.drawSeparator()
	s.drawStopwatch(mu)
}

func (s *Status) drawSeparator() {
	s.screen.SetContent(s.x, 0, '|', nil, tcell.StyleDefault)
	s.x += 1
}

func (s *Status) drawAccuracy(mu *ent.TrainingMutation) {
	accuracy, _ := mu.Accuracy()
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

func (s *Status) drawSpeed(mu *ent.TrainingMutation) {
	speed, _ := mu.Speed()
	for _, r := range fmt.Sprintf("%*.0fw/s", 3, speed) {
		s.screen.SetContent(s.x, 0, r, nil, tcell.StyleDefault)
		s.x += 1
	}
}

func (s *Status) drawProgress(mu *ent.TrainingMutation) {
	progress, _ := mu.Progress()
	for i := 0; i < 10; i++ {
		style := tcell.StyleDefault
		if progress/10.0 > float64(i) {
			style = style.Reverse(true)
		}
		s.screen.SetContent(s.x, 0, ' ', nil, style)
		s.x += 1
	}
}

func (s *Status) drawStopwatch(mu *ent.TrainingMutation) {
	duration, _ := mu.Duration()
	stopwatch, _ := mu.Stopwatch()
	for _, r := range (time.Duration(duration-stopwatch) * time.Second).String() {
		s.screen.SetContent(s.x, 0, r, nil, tcell.StyleDefault)
		s.x += 1
	}
}
