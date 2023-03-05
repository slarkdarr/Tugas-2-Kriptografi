package impl

import (
	"encoding/base64"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/utils"
)

type (
	cipher struct {
		substitution internal.GroupExecutor
		permutation  internal.GroupExecutor
		key          internal.Key
	}
)

func NewCipher(externalKey string) internal.Cipher {
	return &cipher{
		substitution: NewCircularSubsitution(2),
		permutation:  NewCircularPermutation(2),
		key:          NewKey(externalKey),
	}
}

func (c *cipher) blockXOR(a, b []byte) []byte {
	output := make([]byte, 16)
	for i := 0; i < 16; i++ {
		output[i] = a[i] ^ b[i]
	}
	return output
}

func (c *cipher) Encrypt(plaintext string) string {
	c.permutation.ResetCount()
	c.substitution.ResetCount()
	keys := c.GenerateKeys(true)
	blocks := c.GenerateBlocks(plaintext, true)

	encryptedBlocks := make([][]byte, len(blocks))
	for i := 0; i < len(blocks); i++ {
		encryptedBlocks[i] = make([]byte, 16)
	}
	var result []byte
	for i, block := range blocks {
		curr := make([]byte, 16)
		copy(curr, block)
		if i > 0 {
			curr = c.blockXOR(curr, encryptedBlocks[i-1])
		}
		resultBlock := c.Rounds(curr, keys, 0, true)
		copy(encryptedBlocks[i], resultBlock)
		result = append(result, resultBlock...)
	}

	return base64.StdEncoding.EncodeToString(result)
}

func (c *cipher) Decrypt(ciphertext string) string {
	c.permutation.ResetCount()
	c.substitution.ResetCount()
	keys := c.GenerateKeys(false)
	blocks := c.GenerateBlocks(ciphertext, false)

	var result []byte
	for i, block := range blocks {
		resultBlock := c.Rounds(block, keys, 0, false)
		if i > 0 {
			resultBlock = c.blockXOR(resultBlock, blocks[i-1])
		}
		result = append(result, resultBlock...)
	}

	return string(result)
}

func (c *cipher) GenerateKeys(encrypt bool) [][]byte {
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

func (c *cipher) GenerateBlocks(text string, encrypt bool) [][]byte {
	blockSize := 16

	var byteList []byte
	if encrypt {
		byteList = []byte(text)
	} else {
		byteList, _ = base64.StdEncoding.DecodeString(text)
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

func (c *cipher) Rounds(block []byte, keys [][]byte, round int, encrypt bool) []byte {
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

	s1 := c.substitution.Execute(x1)
	s2 := c.substitution.Execute(x2)

	xor := utils.CalculateXor(s1, roundKey[0])
	add := utils.CalculateAddMod32(s2, xor)

	tmp1 := utils.CalculateXor(add, roundKey[1])
	tmp2 := utils.CalculateAddMod32(xor, tmp1)

	p1 := c.permutation.Execute(tmp2)
	p2 := c.permutation.Execute(tmp1)

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
