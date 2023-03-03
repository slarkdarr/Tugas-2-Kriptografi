package impl

import "github.com/slarkdarr/Tugas-2-Kriptografi/internal"

type (
	substitution struct {
		box [][]rune
	}
)

func NewSubstitution() internal.Substitution {
	return &substitution{}
}

func (s substitution) Execute(block []uint8) []uint8 {
	//TODO implement me
	panic("implement me")
}
