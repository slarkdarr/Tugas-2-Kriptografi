package impl

import (
	"bytes"
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
	blocks := c.GenerateBlocks(plaintext)

	var result [][]byte
	for _, block := range blocks {
		result = append(result, c.Rounds(block, keys, 0))
	}

	return string(bytes.Join(result, []byte("")))
}

func (c cipher) Decrypt(ciphertext, externalKey string) string {
	keys := c.GenerateKeys(externalKey, false)
	blocks := c.GenerateBlocks(ciphertext)

	var result [][]byte
	for _, block := range blocks {
		result = append(result, c.Rounds(block, keys, 0))
	}

	return string(bytes.Join(result, []byte("")))
}

func (c cipher) GenerateKeys(externalKey string, encrypt bool) [][]byte {
	return NewKey(externalKey).Generate()
}

func (c cipher) GenerateBlocks(plaintext string) [][]byte {
	byteList := []byte(plaintext)
	blockSize := 16

	remainder := len(byteList) % blockSize
	if remainder != 0 {
		padding := make([]byte, blockSize-remainder)
		byteList = append(byteList, padding...)
	}

	var blocks [][]byte
	for i := 0; i < len(byteList); i += blockSize {
		end := i + blockSize
		if end > len(byteList) {
			end = len(byteList)
		}
		blocks = append(blocks, byteList[i:end])
	}
	return blocks
}

func (c cipher) Rounds(block []byte, keys [][]byte, round int) []byte {
	if round > 16 {
		return block
	}

	roundKey := keys[2*round : 2*round+1]

	x1, x2, x3, x4 := block[0:4], block[4:8], block[8:12], block[12:16]

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

	var newBlock []byte
	newBlock = append(newBlock, newX1...)
	newBlock = append(newBlock, newX2...)
	newBlock = append(newBlock, newX3...)
	newBlock = append(newBlock, newX4...)

	return c.Rounds(newBlock, keys, round+1)
}
