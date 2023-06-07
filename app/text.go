package app

import (
	"github.com/gdamore/tcell/v2"
)

type Text struct {
	screen tcell.Screen

	w int

	h int

	max int

	y int
}

func (t *Text) Draw(text []rune, inputs []bool) {
	textChunks := chunkBy(text, t.max)
	inputChunks := chunkBy(inputs, t.max)
	offset := len(inputs) / t.max
	for i, r := range textChunks[offset] {
		x := i % t.w
		if i > 0 && x == 0 {
			if t.y == t.h {
				break
			}
			t.y += 1
		}

		style := tcell.StyleDefault
		if len(inputChunks) > offset && i < len(inputChunks[offset]) {
			style = style.Bold(true)
			if inputChunks[offset][i] {
				style = style.Foreground(tcell.ColorGreen)
			} else {
				style = style.Foreground(tcell.ColorRed)
			}
			if r == ' ' {
				style = style.Underline(true)
			}
		}

		t.screen.SetContent(x, t.y, r, nil, style)
	}
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
