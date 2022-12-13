package in_memory

import "sync"

type InMemoryStatistic struct {
	sync.RWMutex
	total int
	win   int
	lose  int
	draw  int
}

func (s *InMemoryStatistic) Won() {
	s.Lock()
	defer s.Unlock()
	s.total++
	s.win++
}

func (s *InMemoryStatistic) Lost() {
	s.Lock()
	defer s.Unlock()
	s.total++
	s.lose++
}

func (s *InMemoryStatistic) Draw() {
	s.Lock()
	defer s.Unlock()
	s.total++
	s.draw++
}

func (s *InMemoryStatistic) Get() (int, int, int, int) {
	s.RLock()
	defer s.RUnlock()
	return s.total, s.win, s.lose, s.draw
}
