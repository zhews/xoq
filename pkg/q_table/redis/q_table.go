package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"log"
	"strconv"
)

type QTableRedis struct {
	Client *redis.Client
}

func (qr *QTableRedis) Set(key string, value float64) {
	qr.Client.Set(context.Background(), key, value, 0)
}

func (qr *QTableRedis) Get(key string) float64 {
	result, err := qr.Client.Get(context.Background(), key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Could not retrieve result: ", err)
		}
		return 0
	}
	value, err := strconv.ParseFloat(result, 64)
	if err != nil {
		log.Println("Retrieved an invalid value: ", err)
		return 0
	}
	return value
}
