package util

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func EmitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func FormatDuration(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	switch {
	case h > 0:
		return fmt.Sprintf("%01d:%02d:%02d.%01d", h, m, s, ms/100)
	case m > 0:
		return fmt.Sprintf("%01d:%02d.%01d", m, s, ms/100)
	default:
		return fmt.Sprintf("%01d.%01d", s, ms/100)
	}
}
