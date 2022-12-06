package domain

import "strings"

const (
	height = 3
	width  = 3
)

type Board [height][width]Symbol

func NewEmptyBoard() *Board {
	return &Board{
		{SymbolNone, SymbolNone, SymbolNone},
		{SymbolNone, SymbolNone, SymbolNone},
		{SymbolNone, SymbolNone, SymbolNone},
	}
}

func (b *Board) Hash() string {
	var hash strings.Builder
	for _, row := range b {
		for _, column := range row {
			hash.WriteRune(rune(column))
		}
	}
	return hash.String()
}

func (b *Board) PossibleActions() []Action {
	var actions []Action
	for rowIndex, row := range b {
		for columnIndex, column := range row {
			if column == SymbolNone {
				action := Action{
					Row:    rowIndex,
					Column: columnIndex,
				}
				actions = append(actions, action)
			}
		}
	}
	return actions
}

func (b *Board) IsValidAction(action Action) bool {
	validRow := action.Row >= 0 && action.Row < height
	validColumn := action.Column >= 0 && action.Column < width
	if validRow && validColumn {
		return b[action.Row][action.Column] == SymbolNone
	}
	return false
}

func (b *Board) Winner() Symbol {
	for index, row := range b {
		if row[0] != SymbolNone && row[0] == row[1] && row[1] == row[2] {
			return row[0]
		}
		if b[0][index] != SymbolNone && b[0][index] == b[1][index] && b[1][index] == b[2][index] {
			if b[0][index] != SymbolNone {
				return b[0][index]
			}
		}
	}
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] != SymbolNone {
			return b[0][0]
		}
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		if b[0][2] != SymbolNone {
			return b[0][2]
		}
	}
	return SymbolNone
}

func (b *Board) IsDraw() bool {
	for _, row := range b {
		for _, column := range row {
			if column == SymbolNone {
				return false
			}
		}
	}
	return true
}
