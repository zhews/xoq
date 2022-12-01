package domain

import "math/rand"

type Agent struct {
	QTable QTable
	States []string
}

const (
	experienceRate = 0.2
	learningRate   = 0.2
	gamma          = 0.9
)

func (a *Agent) ChooseAction(board Board) Action {
	if rand.Float64() < experienceRate {
		nextAction := a.chooseRandomAction(board)
		return nextAction
	}
	value, bestAction := a.chooseBestAction(board)
	if value == 0 {
		nextAction := a.chooseRandomAction(board)
		a.updateStates(board, nextAction)
	}
	a.updateStates(board, bestAction)
	return bestAction
}

func (a *Agent) chooseRandomAction(board Board) Action {
	possibleActions := board.PossibleActions()
	randomIndex := rand.Intn(len(possibleActions))
	nextAction := possibleActions[randomIndex]
	return nextAction
}

func (a *Agent) chooseBestAction(board Board) (float64, Action) {
	possibleActions := board.PossibleActions()
	var maxValue float64
	var nextAction Action
	for index, possibleAction := range possibleActions {
		possibleNextBoard := board
		possibleNextBoard[possibleAction.Row][possibleAction.Column] = SymbolAgent
		possibleNextHash := possibleNextBoard.Hash()
		value := a.QTable.Get(possibleNextHash)
		if index == 0 || value > maxValue {
			maxValue = value
			nextAction = possibleAction
		}
	}
	return maxValue, nextAction
}

func (a *Agent) updateStates(board Board, action Action) {
	nextBoard := board
	nextBoard[action.Row][action.Column] = SymbolAgent
	boardHash := nextBoard.Hash()
	a.States = append(a.States, boardHash)
}

func (a *Agent) Reward(amount float64) {
	for latest := len(a.States); latest > 0; latest-- {
		boardHash := a.States[latest]
		currentQValue := a.QTable.Get(boardHash)
		newQValue := currentQValue + learningRate*(amount*gamma-currentQValue)
		a.QTable.Set(boardHash, newQValue)
		amount = newQValue
	}
}
