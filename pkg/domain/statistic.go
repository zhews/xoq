package domain

import "sync"

type Statistic struct {
	sync.RWMutex
	total int
	win   int
	lose  int
	draw  int
}

func (s *Statistic) Won() {
	s.Lock()
	defer s.Unlock()
	s.total++
	s.win++
}

func (s *Statistic) Lost() {
	s.Lock()
	defer s.Unlock()
	s.total++
	s.lose++
}

func (s *Statistic) Draw() {
	s.Lock()
	defer s.Unlock()
	s.total++
	s.draw++
}

func (s *Statistic) Get() (int, int, int, int) {
	s.RLock()
	defer s.RUnlock()
	return s.total, s.win, s.lose, s.draw
}
