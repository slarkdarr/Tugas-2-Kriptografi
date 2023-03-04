package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
)

type (
	substitution struct {
		box [][]rune
	}
)

func NewSubstitution() internal.Substitution {
	return &substitution{}
}

func (s substitution) Execute(chunk uint32) uint32 {
	//TODO implement me
	panic("implement me")
}
