package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"xoq/pkg/handler"
	qTable "xoq/pkg/q_table/in_memory"
	statistic "xoq/pkg/statistic/in_memory"
)

func RunHTTPServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5173"}))
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
	log.Fatal(app.Listen(":8080"))
}
