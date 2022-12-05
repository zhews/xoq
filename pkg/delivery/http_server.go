package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"log"
	"xoq/pkg/domain"
	"xoq/pkg/handler"
)

func RunHTTPServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5173"}))
	app.Use("/game", handler.UpgradeToWebsocket)
	qTable := domain.NewQTable()
	statistic := &domain.Statistic{}
	gameHandler := handler.GameHandler{
		QTable:    qTable,
		Statistic: statistic,
	}
	app.Get("/game", websocket.New(gameHandler.RunGame))
	statisticHandler := handler.StatisticHandler{
		Statistic: statistic,
	}
	app.Get("/statistic", statisticHandler.Current)
	log.Fatal(app.Listen(":8080"))
}
