package in_memory

import "sync"

type InMemoryQTable struct {
	sync.RWMutex
	qValues map[string]float64
}

func (qt *InMemoryQTable) Set(key string, value float64) {
	qt.Lock()
	defer qt.Unlock()
	qt.qValues[key] = value
}

func (qt *InMemoryQTable) Get(key string) float64 {
	qt.RLock()
	defer qt.RUnlock()
	return qt.qValues[key]
}

func NewQTable() *InMemoryQTable {
	return &InMemoryQTable{
		qValues: make(map[string]float64),
	}
}
