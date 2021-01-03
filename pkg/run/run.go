package run

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
)

type Run struct {
	sync.RWMutex

	ID             string
	Name           string
	Category       string
	segments       []*Segment
	started, ended time.Time
	paused         time.Duration
}

func Load(filename string) (*Run, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var file struct {
		ID       string
		Name     string
		Category string
		Segments []*Segment
	}
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, err
	}
	return &Run{
		ID:       file.ID,
		Name:     file.Name,
		Category: file.Category,
		segments: file.Segments,
	}, nil
}

func (r *Run) Save(filename string) error {
	segments := make([]*Segment, 0)
	for _, segment := range r.segments {
		s := &Segment{
			Name: segment.Name,
		}
		if segment.SplitTime < segment.BestSplitTime || segment.BestSplitTime == 0 {
			s.BestSplitTime = segment.SplitTime
		}
		segments = append(segments, s)
	}
	data, err := json.MarshalIndent(struct {
		ID       string
		Name     string
		Category string
		Segments []*Segment
	}{
		ID:       r.ID,
		Name:     r.Name,
		Category: r.Category,
		Segments: segments,
	}, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (r *Run) Start() {
	r.Lock()
	defer r.Unlock()

	if !r.started.IsZero() {
		return
	}
	r.started = time.Now().Add(-15 * time.Minute)
}

func (r *Run) Stop() {
	r.Lock()
	defer r.Unlock()

	if !r.ended.IsZero() {
		return
	}
	r.ended = time.Now()
}

func (r *Run) Split() {
	r.Lock()
	defer r.Unlock()

	for i, s := range r.segments {
		if len(r.segments)-1 == i {
			if r.ended.IsZero() {
				r.ended = time.Now()
			}
		}
		if s.SplitTime == 0 {
			s.SplitTime = time.Since(r.started)
			break
		}
	}
}

func (r *Run) Reset() {
	r.Lock()
	defer r.Unlock()

	r.started = time.Time{}
	r.ended = time.Time{}
	for i := range r.segments {
		r.segments[i].SplitTime = 0
	}
}

func (r *Run) Segments() []*Segment {
	r.RLock()
	defer r.RUnlock()

	return r.segments
}

func (r *Run) Running() bool {
	r.RLock()
	defer r.RUnlock()

	return !r.started.IsZero()
}

func (r *Run) CurrentTime() time.Duration {
	r.RLock()
	defer r.RUnlock()

	if r.started.IsZero() {
		return 0
	}
	if r.ended.IsZero() {
		return time.Since(r.started)
	}
	return r.ended.Sub(r.started)
}
