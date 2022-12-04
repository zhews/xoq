package dto

import "xoq/pkg/domain"

const (
	ResponseTypeError  = "error"
	ResponseTypeWinner = "winner"
	ResponseTypeDraw   = "draw"
	ResponseTypeBoard  = "board"
)

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseWinner struct {
	Symbol domain.Symbol `json:"symbol"`
}
