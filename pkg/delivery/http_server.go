package delivery

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"xoq/pkg/config"
	"xoq/pkg/handler"
	qTable "xoq/pkg/q_table/in_memory"
	statistic "xoq/pkg/statistic/in_memory"
)

func RunHTTPServer() {
	xoqConfig, err := config.ParseFromEnvironmentalVariables()
	if err != nil {
		log.Fatal(xoqConfig)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: xoqConfig.CorsAllowOrigins}))
	app.Use("/game", handler.UpgradeToWebsocket)
	inMemoryQTable := qTable.NewQTable()
	inMemoryStatistic := &statistic.InMemoryStatistic{}
	gameHandler := handler.GameHandler{
		QTable:    inMemoryQTable,
		Statistic: inMemoryStatistic,
	}
	app.Get("/game", websocket.New(gameHandler.RunGame))
	statisticHandler := handler.StatisticHandler{
		Statistic: inMemoryStatistic,
	}
	app.Get("/statistic", statisticHandler.Current)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", xoqConfig.Port)))
}
