package tui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"

	"github.com/ChrisRx/splits/pkg/run"
	"github.com/ChrisRx/splits/pkg/util"
)

func DrawSegments(scr tcell.Screen, style tcell.Style, segments []*run.Segment) {
	w, _ := scr.Size()
	for i, segment := range segments {
		elapsed := util.FormatDuration(segment.SplitTime)
		if segment.SplitTime == 0 {
			elapsed = "-"
		}
		delta := segment.Delta()
		padding2 := strings.Repeat(" ", 10-(len(elapsed)%10))
		padding := strings.Repeat(" ", w-len(segment.Name)-len(delta)-len(elapsed)-(10-len(elapsed)))
		util.EmitStr(scr, 0, 5+i, style, fmt.Sprintf("%s%s%s%s%s", segment.Name, padding, delta, padding2, elapsed))
	}
}

func DrawTitle(s tcell.Screen, title, category string) {
	w, _ := s.Size()
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(0x2f3542)).
		Foreground(tcell.ColorWhite).
		Bold(true)
	headerStyle := tcell.StyleDefault.
		Background(tcell.NewHexColor(0x2f3542)).
		Foreground(tcell.NewHexColor(0xa4b0be))
	util.EmitStr(s, w/2-(len(title)/2), 0, style, title)
	util.EmitStr(s, w/2-(len(category)/2), 1, style, category)
	splitHeader := "Delta      Time"
	util.EmitStr(s, w-(len(splitHeader)), 3, headerStyle, splitHeader)
	util.EmitStr(s, 0, 4, headerStyle, strings.Repeat("─", w))
}

func DrawTimer(s tcell.Screen, style tcell.Style, ct string) {
	var sb strings.Builder
	for i := 0; i < 4; i++ {
		for _, c := range ct {
			char, ok := util.BigCharMap[c]
			if !ok {
				continue
			}
			sb.WriteString(char[i])
		}
		line := sb.String()
		w, _ := s.Size()
		var l int
		for _, c := range line {
			l += runewidth.RuneWidth(c)
		}
		util.EmitStr(s, w-l, 15+i, style, line)
		sb.Reset()
	}
}
