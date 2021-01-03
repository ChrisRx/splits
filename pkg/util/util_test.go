package util

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestFormatDuration(t *testing.T) {
	cases := []struct {
		d        time.Duration
		expected string
	}{
		{
			5*time.Hour + 4*time.Minute + 205*time.Millisecond,
			"5:04:00.2",
		},
		{
			4*time.Minute + 15*time.Second + 205*time.Millisecond,
			"4:15.2",
		},
		{
			15*time.Second + 847*time.Millisecond,
			"15.8",
		},
		{
			121*time.Second + 847*time.Millisecond,
			"2:01.8",
		},
	}

	for _, tc := range cases {
		result := FormatDuration(tc.d)
		if diff := cmp.Diff(result, tc.expected); diff != "" {
			t.Fatalf("%v", diff)
		}
	}
}
