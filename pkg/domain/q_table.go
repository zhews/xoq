package domain

type QTable interface {
	Set(key string, value float64)
	Get(key string) float64
}
