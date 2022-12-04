package handler

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"xoq/pkg/domain"
)

type GameHandler struct {
	QTable *domain.QTable
}

func (h *GameHandler) RunGame(conn *websocket.Conn) {
	defer conn.Close()
	agent := domain.Agent{QTable: h.QTable}
	board := domain.NewEmptyBoard()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var playerAction domain.Action
		if err = json.Unmarshal(message, &playerAction); err != nil {
			conn.WriteJSON(Response{Type: ResponseTypeError, Data: ResponseMessage{Message: "Invalid message!"}})
			continue
		}
		if !board.IsValidAction(playerAction) {
			conn.WriteJSON(Response{Type: ResponseTypeError, Data: ResponseMessage{Message: "Invalid action!"}})
			continue
		}
		board[playerAction.Row][playerAction.Column] = domain.SymbolPlayer
		if gameIsFinished(conn, board, agent) {
			break
		}
		agentAction := agent.ChooseAction(*board)
		board[agentAction.Row][agentAction.Column] = domain.SymbolAgent
		if gameIsFinished(conn, board, agent) {
			break
		}
		conn.WriteJSON(Response{Type: ResponseTypeBoard, Data: board})
	}
}

func gameIsFinished(conn *websocket.Conn, board *domain.Board, agent domain.Agent) bool {
	winner := board.Winner()
	if winner != domain.SymbolNone {
		if winner == domain.SymbolAgent {
			agent.Reward(1)
		} else {
			agent.Reward(0)
		}
		conn.WriteJSON(Response{Type: ResponseTypeWinner, Data: ResponseWinner{Symbol: winner}})
		return true
	} else {
		if board.IsDraw() {
			agent.Reward(0.3)
			conn.WriteJSON(Response{Type: ResponseTypeDraw, Data: nil})
			return true
		}
	}
	return false
}
