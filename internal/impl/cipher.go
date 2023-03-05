package impl

import (
	"encoding/hex"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/utils"
)

type (
	cipher struct {
		substitution internal.Substitution
		permutation  internal.Permutation
		key          internal.Key
	}
)

func NewCipher(externalKey string) internal.Cipher {
	return &cipher{
		substitution: NewSubstitution(),
		permutation:  NewPermutation(),
		key:          NewDummyKey(),
	}
}

func (c cipher) Encrypt(plaintext string) string {
	keys := c.GenerateKeys(true)
	blocks := c.GenerateBlocks(plaintext, true)

	var result []byte
	for _, block := range blocks {
		result = append(result, c.Rounds(block, keys, 0, true)...)
	}

	return hex.EncodeToString(result)
}

func (c cipher) Decrypt(ciphertext string) string {
	keys := c.GenerateKeys(false)
	blocks := c.GenerateBlocks(ciphertext, false)

	var result []byte
	for _, block := range blocks {
		result = append(result, c.Rounds(block, keys, 0, false)...)
	}

	return string(result)
}

func (c cipher) GenerateKeys(encrypt bool) [][]byte {
	result := c.key.Generate()
	if encrypt {
		return result
	}

	var keyList [][]byte
	for i := len(result) - 1; i > 0; i -= 2 {
		keyList = append(keyList, result[i-1], result[i])
	}

	return keyList
}

func (c cipher) GenerateBlocks(text string, encrypt bool) [][]byte {
	blockSize := 16

	var byteList []byte
	if encrypt {
		byteList = []byte(text)
	} else {
		byteList, _ = hex.DecodeString(text)
	}

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

func (c cipher) Rounds(block []byte, keys [][]byte, round int, encrypt bool) []byte {
	if round >= 16 {
		return block
	}

	roundKey := keys[2*round : 2*round+2]

	var x1, x2, x3, x4 []byte
	if encrypt {
		x1, x2, x3, x4 = block[0:4], block[4:8], block[8:12], block[12:16]
	} else {
		x1, x2, x3, x4 = block[8:12], block[12:16], block[0:4], block[4:8]
	}

	s1 := c.substitution.Execute(x1, encrypt)
	s2 := c.substitution.Execute(x2, encrypt)

	xor := utils.CalculateXor(s1, roundKey[0])
	add := utils.CalculateAddMod32(s2, xor)

	tmp1 := utils.CalculateXor(add, roundKey[1])
	tmp2 := utils.CalculateAddMod32(xor, tmp1)

	p1 := c.permutation.Execute(tmp2, encrypt)
	p2 := c.permutation.Execute(tmp1, encrypt)

	var newX1, newX2, newX3, newX4 []byte
	if encrypt {
		newX1, newX2, newX3, newX4 = utils.CalculateXor(p1, x3), utils.CalculateXor(p2, x4), x1, x2
	} else {
		newX1, newX2, newX3, newX4 = x1, x2, utils.CalculateXor(p1, x3), utils.CalculateXor(p2, x4)
	}

	var newBlock []byte
	newBlock = append(newBlock, newX1...)
	newBlock = append(newBlock, newX2...)
	newBlock = append(newBlock, newX3...)
	newBlock = append(newBlock, newX4...)

	return c.Rounds(newBlock, keys, round+1, encrypt)
}
