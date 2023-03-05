package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"math/rand"
	"sync"
)

type (
	substitution struct {
		box [256]byte
	}
)

var once sync.Once
var s rand.Source

func NewSubstitution() internal.Executor {
	once.Do(func() {
		s = rand.NewSource(202303051750)
	})
	var sbox [256]byte
	for i := 0; i < 256; i++ {
		sbox[i] = byte(i)
	}

	Intn := func(i int) int {
		num64 := s.Int63()
		num := int(num64)
		return num % i
	}

	for k := 0; k < 16; k++ {
		for i := 255; i > 0; i-- {

			j := Intn(i + 1)
			sbox[i], sbox[j] = sbox[j], sbox[i]
		}
	}
	return &substitution{
		box: sbox,
	}
}

func (s *substitution) Execute(chunk []byte) []byte {
	validatedChunk := s.validate(chunk)
	return s.forward(validatedChunk)
}

func (s *substitution) validate(chunk []byte) []byte {
	result := make([]byte, 4)
	for i := 0; i < 4; i++ {
		if i+1 > len(chunk) {
			result[i] = 0
		} else {
			result[i] = chunk[i]
		}
	}
	return result
}

func (s *substitution) forward(chunk []byte) []byte {
	output := make([]byte, 4)
	for i, b := range chunk {
		output[i] = s.box[b]
	}
	return output
}

func (s *substitution) reverse(chunk []byte) []byte {
	output := make([]byte, 4)
	for i, b := range chunk {
		for j, value := range s.box {
			if value == b {
				output[i] = byte(j)
				break
			}
		}
	}
	return output[0:4]
}
