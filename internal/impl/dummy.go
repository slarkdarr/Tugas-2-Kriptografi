package impl

import "github.com/slarkdarr/Tugas-2-Kriptografi/internal"

type (
	dummy struct{}
)

func (d dummy) Generate() [][]byte {
	slice := make([][]byte, 32)

	for i := range slice {
		slice[i] = make([]byte, 4)
	}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			slice[i][j] = byte(i*len(slice) + j)
		}
	}
	return slice
}

func NewDummyKey() internal.Key {
	return &dummy{}
}
