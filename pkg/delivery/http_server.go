package delivery

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"xoq/pkg/handler"
	qTable "xoq/pkg/q_table/redis"
	statistic "xoq/pkg/statistic/redis"
)

func RunHTTPServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5173"}))
	app.Use("/game", handler.UpgradeToWebsocket)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisQTable := &qTable.QTableRedis{Client: rdb}
	redisStatistic := &statistic.StatisticRedis{Client: rdb}
	gameHandler := handler.GameHandler{
		QTable:    redisQTable,
		Statistic: redisStatistic,
	}
	app.Get("/game", websocket.New(gameHandler.RunGame))
	statisticHandler := handler.StatisticHandler{
		Statistic: redisStatistic,
	}
	app.Get("/statistic", statisticHandler.Current)
	log.Fatal(app.Listen(":8080"))
}
