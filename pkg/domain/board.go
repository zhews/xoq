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
