package tui

import (
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"

	"github.com/ChrisRx/splits/pkg/run"
	"github.com/ChrisRx/splits/pkg/util"
)

var (
	AheadStyle = tcell.StyleDefault.
			Background(tcell.NewHexColor(0x2f3542)).
			Foreground(tcell.ColorGreen)
	BehindStyle = tcell.StyleDefault.
			Background(tcell.NewHexColor(0x2f3542)).
			Foreground(tcell.ColorRed)
	GoldStyle = tcell.StyleDefault.
			Background(tcell.NewHexColor(0x2f3542)).
			Foreground(tcell.ColorGold)
)

func DrawSegments(scr tcell.Screen, style tcell.Style, segments []*run.Segment) {
	w, _ := scr.Size()
	var lastSegment, lastBest time.Duration
	for i, segment := range segments {
		var delta, elapsed string
		deltaStyle := style
		switch {
		case segment.SplitTime == 0:
			delta = "-"
			elapsed = "-"
		case segment.BestSplitTime == 0:
			delta = "-"
			elapsed = util.FormatDuration(segment.SplitTime)
		case segment.BestSplitTime != 0:
			segmentDiff := segment.SplitTime - lastSegment
			lastSegment = segment.SplitTime
			bestDiff := segment.BestSplitTime - lastBest
			lastBest = segment.BestSplitTime
			elapsed = util.FormatDuration(segment.SplitTime)
			saved := segmentDiff - bestDiff
			saved = saved.Truncate(100 * time.Millisecond)
			if saved < 0 {
				saved *= -1
			}
			delta = util.FormatDuration(saved)
			dt := segment.SplitTime - segment.BestSplitTime
			switch {
			case dt > 0:
				deltaStyle = BehindStyle
				delta = "+" + delta
			default:
				deltaStyle = AheadStyle
				delta = "-" + delta
			}
			switch {
			case saved == 0:
				deltaStyle = style
			case segmentDiff < bestDiff:
				deltaStyle = GoldStyle
			}
		}
		padding := w - len(segment.Name) - len(delta) - len(elapsed) - (10 - len(elapsed))
		padding2 := 10 - (len(elapsed) % 10)
		util.EmitStr(scr, 0, 5+i, style, segment.Name)
		util.EmitStr(scr, len(segment.Name)+padding, 5+i, deltaStyle, delta)
		util.EmitStr(scr, len(segment.Name)+padding+len(delta)+padding2, 5+i, style, elapsed)
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
