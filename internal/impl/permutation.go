package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
)

type (
	permutation struct {
		box [32]int
	}
)

func NewPermutation() internal.Permutation {
	return &permutation{
		box: [32]int{
			31, 24, 19, 10, 12, 4, 21, 27,
			20, 14, 28, 17, 30, 9, 1, 7,
			23, 16, 8, 13, 25, 2, 29, 18,
			22, 11, 3, 26, 6, 5, 15, 0,
		},
	}
}

func (p *permutation) Execute(chunk []byte, encrypt bool) []byte {
	validatedChunk := p.validate(chunk)
	if !encrypt {
		return p.reverse(validatedChunk)
	}
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
