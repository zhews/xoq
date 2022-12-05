package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"xoq/pkg/domain"
	"xoq/pkg/handler"
)

func RunHTTPServer() {
	app := fiber.New()
	app.Use("/game", handler.UpgradeToWebsocket)
	qTable := domain.NewQTable()
	statistic := &domain.Statistic{}
	gameHandler := handler.GameHandler{
		QTable:    qTable,
		Statistic: statistic,
	}
	app.Get("/game", websocket.New(gameHandler.RunGame))
	log.Fatal(app.Listen(":8080"))
}
