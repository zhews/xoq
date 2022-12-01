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
