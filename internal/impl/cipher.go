package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/utils"
)

type (
	cipher struct {
		substitution internal.Substitution
		permutation  internal.Permutation
	}
)

func NewCipher() internal.Cipher {
	return &cipher{
		substitution: NewSubstitution(),
		permutation:  NewPermutation(),
	}
}

func (c cipher) Encrypt(plaintext, externalKey string) string {
	keys := c.GenerateKeys(externalKey, true)
	blocks := c.GenerateBlocks(plaintext, true)

	var result [][]uint32
	for _, block := range blocks {
		result = append(result, c.Rounds(block, keys, 0))
	}

	panic("implement me")
}

func (c cipher) Decrypt(ciphertext, externalKey string) string {
	keys := c.GenerateKeys(externalKey, false)
	blocks := c.GenerateBlocks(ciphertext, false)

	var result [][]uint32
	for _, block := range blocks {
		result = append(result, c.Rounds(block, keys, 0))
	}

	panic("implement me")
}

func (c cipher) GenerateKeys(externalKey string, encrypt bool) []uint32 {
	return NewKey(externalKey).Generate()
}

func (c cipher) GenerateBlocks(plaintext string, encrypt bool) [][]uint32 {
	var blocks [][]uint32
	return blocks
}

func (c cipher) Rounds(block []uint32, keys []uint32, round int) []uint32 {
	if round > 16 {
		return block
	}

	roundKey := keys[2*round : 2*round+1]

	x1, x2, x3, x4 := block[0], block[1], block[2], block[3]

	s1 := c.substitution.Execute(x1)
	s2 := c.substitution.Execute(x2)

	xor := utils.CalculateXor(s1, roundKey[0])
	add := utils.CalculateAddMod32(s2, xor)

	tmp1 := utils.CalculateXor(add, roundKey[1])
	tmp2 := utils.CalculateAddMod32(xor, tmp1)

	p1 := c.permutation.Execute(tmp2)
	p2 := c.permutation.Execute(tmp1)

	newX1 := utils.CalculateXor(p1, x3)
	newX2 := utils.CalculateXor(p2, x4)
	newX3 := x1
	newX4 := x2

	var newBlock []uint32
	newBlock = append(newBlock, newX1, newX2, newX3, newX4)

	return c.Rounds(newBlock, keys, round+1)
}
