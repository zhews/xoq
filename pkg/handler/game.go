package handler

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"xoq/pkg/domain"
	"xoq/pkg/handler/dto"
)

type GameHandler struct {
	QTable    *domain.QTable
	Statistic *domain.Statistic
}

func (gh *GameHandler) RunGame(conn *websocket.Conn) {
	defer conn.Close()
	agent := domain.Agent{QTable: gh.QTable}
	board := domain.NewEmptyBoard()
	conn.WriteJSON(dto.Response{Type: dto.ResponseTypeBoard, Data: board})
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var playerAction domain.Action
		if err = json.Unmarshal(message, &playerAction); err != nil {
			conn.WriteJSON(dto.Response{Type: dto.ResponseTypeError, Data: dto.ResponseMessage{Message: "Invalid message!"}})
			continue
		}
		if !board.IsValidAction(playerAction) {
			conn.WriteJSON(dto.Response{Type: dto.ResponseTypeError, Data: dto.ResponseMessage{Message: "Invalid action!"}})
			continue
		}
		board[playerAction.Row][playerAction.Column] = domain.SymbolPlayer
		if gh.gameIsFinished(conn, board, agent) {
			conn.WriteJSON(dto.Response{Type: dto.ResponseTypeBoard, Data: board})
			break
		}
		agentAction := agent.ChooseAction(*board)
		board[agentAction.Row][agentAction.Column] = domain.SymbolAgent
		if gh.gameIsFinished(conn, board, agent) {
			conn.WriteJSON(dto.Response{Type: dto.ResponseTypeBoard, Data: board})
			break
		}
		conn.WriteJSON(dto.Response{Type: dto.ResponseTypeBoard, Data: board})
	}
}

func (gh *GameHandler) gameIsFinished(conn *websocket.Conn, board *domain.Board, agent domain.Agent) bool {
	winner := board.Winner()
	if winner != domain.SymbolNone {
		if winner == domain.SymbolAgent {
			agent.Reward(1)
			gh.Statistic.Won()
		} else {
			agent.Reward(0)
			gh.Statistic.Lost()
		}
		conn.WriteJSON(dto.Response{Type: dto.ResponseTypeWinner, Data: dto.ResponseWinner{Symbol: winner}})
		return true
	} else {
		if board.IsDraw() {
			agent.Reward(0.3)
			gh.Statistic.Draw()
			conn.WriteJSON(dto.Response{Type: dto.ResponseTypeDraw, Data: nil})
			return true
		}
	}
	return false
}
