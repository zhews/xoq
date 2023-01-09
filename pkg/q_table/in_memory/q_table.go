package in_memory

import (
	"encoding/json"
	"os"
	"sync"
)

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

func (qt *InMemoryQTable) WriteToDisk(filename string) {
	qValues, err := json.Marshal(qt.qValues)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, qValues, 0755)
	if err != nil {
		panic(err)
	}
}

func NewQTable() *InMemoryQTable {
	return &InMemoryQTable{
		qValues: make(map[string]float64),
	}
}

func LoadFromDisk() *InMemoryQTable {
	var qValues map[string]float64
	policy, err := os.ReadFile("policy.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(policy, &qValues)
	if err != nil {
		panic(err)
	}
	return &InMemoryQTable{
		qValues: qValues,
	}
}
