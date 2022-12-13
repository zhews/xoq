package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"log"
	"strconv"
)

type StatisticRedis struct {
	Client *redis.Client
}

const (
	KeyTotal  = "total"
	KeyWins   = "wins"
	KeyLosses = "loss"
	KeyDraws  = "draw"
)

func (sr *StatisticRedis) Won() {
	sr.Client.Incr(context.Background(), KeyTotal)
	sr.Client.Incr(context.Background(), KeyWins)
}

func (sr *StatisticRedis) Lost() {
	sr.Client.Incr(context.Background(), KeyTotal)
	sr.Client.Incr(context.Background(), KeyLosses)
}

func (sr *StatisticRedis) Draw() {
	sr.Client.Incr(context.Background(), KeyTotal)
	sr.Client.Incr(context.Background(), KeyDraws)
}

func (sr *StatisticRedis) Get() (int, int, int, int) {
	totalResult, err := sr.Client.Get(context.Background(), KeyTotal).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Could not get total result: ", err)
			return 0, 0, 0, 0
		}
		totalResult = "0"
	}
	total, err := strconv.Atoi(totalResult)
	if err != nil {
		log.Println("Received an invalid total value: ", err)
		return 0, 0, 0, 0
	}
	winsResult, err := sr.Client.Get(context.Background(), KeyWins).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Could not get win result: ", err)
			return 0, 0, 0, 0
		}
		winsResult = "0"
	}
	wins, err := strconv.Atoi(winsResult)
	if err != nil {
		log.Println("Received an invalid wins value: ", err)
		return 0, 0, 0, 0
	}
	lossesResult, err := sr.Client.Get(context.Background(), KeyLosses).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Could not get win result: ", err)
			return 0, 0, 0, 0
		}
		lossesResult = "0"
	}
	losses, err := strconv.Atoi(lossesResult)
	if err != nil {
		log.Println("Received an invalid wins value: ", err)
		return 0, 0, 0, 0
	}
	drawsResult, err := sr.Client.Get(context.Background(), KeyLosses).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println("Could not get win result: ", err)
			return 0, 0, 0, 0
		}
		drawsResult = "0"
	}
	draws, err := strconv.Atoi(drawsResult)
	if err != nil {
		log.Println("Received an invalid wins value: ", err)
		return 0, 0, 0, 0
	}
	return total, wins, losses, draws
}
