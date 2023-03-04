package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/utils"
)

type (
	key struct {
		externalKey string
	}
)

func NewKey(externalKey string) internal.Key {
	return &key{externalKey}
}

func (k key) Generate() [][]byte {
	//TODO implement me
	panic("implement me")
}

func (k key) KeyWhitening() [][]byte {
	keyByteList := []byte(k.externalKey)

	firstKeyConstant := []byte{45, 233, 169, 143}
	secondKeyConstant := []byte{181, 252, 21, 242}

	var intermediateKey [][]byte
	for i := 0; i < 16; i += 4 {
		byteList := keyByteList[i : i+4]
		intermediateKey = append(intermediateKey, byteList)
	}

	for i := 0; i < 4; i++ {
		intermediateByte := utils.CalculateXor(intermediateKey[i], firstKeyConstant)
		intermediateByte = utils.CalculateXor(intermediateByte, secondKeyConstant)

		intermediateKey[i] = intermediateByte
	}

	return intermediateKey
}

func (k key) KeySchedule() [][]byte {
	return
}
