package util

import "strings"

var BigCharMap = map[rune][4]string{
	'0': {
		"▄▀▀▄ ",
		"█  █ ",
		"█  █ ",
		" ▀▀  ",
	},
	'1': {
		" ▄█  ",
		"  █  ",
		"  █  ",
		" ▀▀▀ ",
	},
	'2': {
		"▄▀▀▄ ",
		"  ▄▀ ",
		"▄▀   ",
		"▀▀▀▀ ",
	},
	'3': {
		"▄▀▀▄ ",
		"  ▄▀ ",
		"▄  █ ",
		" ▀▀  ",
	},
	'4': {
		"▄  █ ",
		"█▄▄█ ",
		"   █ ",
		"   ▀ ",
	},
	'5': {
		"█▀▀▀ ",
		"█▄▄  ",
		"   █ ",
		"▀▀▀  ",
	},
	'6': {
		"▄▀▀  ",
		"█▄▄  ",
		"█  █ ",
		" ▀▀  ",
	},
	'7': {
		"▀▀▀█ ",
		"  ▐▌ ",
		"  █  ",
		"  ▀  ",
	},
	'8': {
		"▄▀▀▄ ",
		"▀▄▄▀ ",
		"█  █ ",
		" ▀▀  ",
	},
	'9': {
		"▄▀▀▄ ",
		"▀▄▄█ ",
		"   █ ",
		" ▀▀  ",
	},
	'.': {
		"  ",
		"  ",
		"  ",
		"▀ ",
	},
	':': {
		"  ",
		"▀ ",
		"▀ ",
		"  ",
	},
}

func BigDigits(s string) string {
	var sb strings.Builder
	for i := 0; i < 4; i++ {
		sb.WriteString("\r")
		for _, c := range s {
			char, ok := BigCharMap[c]
			if !ok {
				continue
			}
			sb.WriteString(char[i])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
