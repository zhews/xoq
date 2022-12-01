package domain

import "sync"

type QTable struct {
	sync.RWMutex
	qValues map[string]float64
}

func (qt *QTable) Set(key string, value float64) {
	qt.Lock()
	defer qt.Unlock()
	qt.qValues[key] = value
}

func (qt *QTable) Get(key string) float64 {
	qt.RLock()
	defer qt.RUnlock()
	return qt.qValues[key]
}
