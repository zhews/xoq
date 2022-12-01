package domain

import "strings"

type Board [3][3]Symbol

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
