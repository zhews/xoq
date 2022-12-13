package handler

import (
	"github.com/gofiber/fiber/v2"
	"xoq/pkg/domain"
	"xoq/pkg/handler/dto"
)

type StatisticHandler struct {
	Statistic domain.Statistic
}

func (sh *StatisticHandler) Current(ctx *fiber.Ctx) error {
	total, win, loss, draw := sh.Statistic.Get()
	response := dto.ResponseStatistic{
		Total: total,
		Win:   win,
		Lose:  loss,
		Draw:  draw,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
