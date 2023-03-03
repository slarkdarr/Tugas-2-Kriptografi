package impl

import "github.com/slarkdarr/Tugas-2-Kriptografi/internal"

type (
	permutation struct {
		box []rune
	}
)

func NewPermutation() internal.Permutation {
	return &permutation{}
}

func (p permutation) Execute(block []uint8) []uint8 {
	//TODO implement me
	panic("implement me")
}
