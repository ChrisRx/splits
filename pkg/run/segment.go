package run

import (
	"fmt"
	"time"

	"github.com/ChrisRx/splits/pkg/util"
)

type Segment struct {
	Name          string
	SplitTime     time.Duration `json:"-"`
	BestSplitTime time.Duration
}

func (s *Segment) Delta() string {
	if s.SplitTime == 0 {
		return "-"
	}
	if s.BestSplitTime == 0 {
		return "-"
	}
	diff := s.SplitTime - s.BestSplitTime
	if diff > 0 {
		return fmt.Sprintf("+%s", util.FormatDuration(diff))
	}
	diff *= -1
	return fmt.Sprintf("-%s", util.FormatDuration(diff))
}
