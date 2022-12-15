package main

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"xoq/pkg/config"
	"xoq/pkg/handler"
	qTable "xoq/pkg/q_table/redis"
	statistic "xoq/pkg/statistic/redis"
)

func main() {
	xoqConfig, err := config.ParseFromEnvironmentalVariables()
	if err != nil {
		log.Fatal(xoqConfig)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     xoqConfig.RedisConfig.Host,
		Username: xoqConfig.RedisConfig.User,
		Password: xoqConfig.RedisConfig.Password,
	})
	redisQTable := &qTable.QTableRedis{Client: redisClient}
	redisStatistic := &statistic.StatisticRedis{Client: redisClient}
	gameHandler := handler.GameHandler{
		QTable:    redisQTable,
		Statistic: redisStatistic,
	}
	statisticHandler := handler.StatisticHandler{
		Statistic: redisStatistic,
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: xoqConfig.CorsAllowOrigins}))
	app.Use("/game", handler.UpgradeToWebsocket)
	app.Get("/game", websocket.New(gameHandler.RunGame))
	app.Get("/statistic", statisticHandler.Current)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", xoqConfig.Port)))
}
