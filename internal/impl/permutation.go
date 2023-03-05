package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"math/rand"
)

type (
	permutation struct {
		box [32]int
	}
)

func NewPermutation() internal.Executor {
	once.Do(func() {
		s = rand.NewSource(202303051750)
	})
	var pbox [32]int
	for i := 0; i < 32; i++ {
		pbox[i] = i
	}

	Intn := func(i int) int {
		num64 := s.Int63()
		num := int(num64)
		return num % i
	}

	for k := 0; k < 16; k++ {
		for i := 31; i > 0; i-- {
			j := Intn(i + 1)
			pbox[i], pbox[j] = pbox[j], pbox[i]
		}
	}
	return &permutation{
		box: pbox,
	}
}

func (p *permutation) Execute(chunk []byte) []byte {
	validatedChunk := p.validate(chunk)
	return p.forward(validatedChunk)
}

func (p *permutation) forward(chunk []byte) []byte {
	output := make([]byte, 4)
	for i := 0; i < 32; i++ {
		inputIndex := p.box[i] >> 3
		inputBit := p.box[i] & 0x07
		outputIndex := i >> 3
		outputBit := uint(7 - (i & 0x07))
		if (chunk[inputIndex] & (1 << inputBit)) != 0 {
			output[outputIndex] |= (1 << outputBit)
		}
	}
	return output
}

func (p *permutation) reverse(chunk []byte) []byte {
	output := make([]byte, 4)
	for i := 0; i < 32; i++ {
		inputIndex := i >> 3
		inputBit := uint(7 - (i & 0x07))
		outputIndex := p.box[i] >> 3
		outputBit := p.box[i] & 0x07
		if (chunk[inputIndex] & (1 << inputBit)) != 0 {
			output[outputIndex] |= (1 << outputBit)
		}
	}
	return output[0:4]
}

func (p *permutation) validate(chunk []byte) []byte {
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
