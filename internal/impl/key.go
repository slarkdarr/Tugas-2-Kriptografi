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
	intermediateKey := k.KeyWhitening()
	keySchedule := k.KeySchedule(intermediateKey)
	sBox := k.GenerateSBox(keySchedule)

	return sBox
}

// KeyWhitening : XOR with constants
func (k key) KeyWhitening() [][]byte {
	keyByteList := []byte(k.externalKey)

	// Whitening constants
	firstKeyConstant := []byte{45, 233, 169, 143}
	secondKeyConstant := []byte{181, 252, 21, 242}

	var intermediateKey [][]byte
	// Convert keyByteList to 2D array (4 x 4 bytes)
	for i := 0; i < 16; i += 4 {
		byteList := keyByteList[i : i+4]
		intermediateKey = append(intermediateKey, byteList)
	}

	// Whitening procedure
	// XOR each 4-byte in intermediateKey array with two constants
	for i := 0; i < 4; i++ {
		intermediateByte := utils.CalculateXor(intermediateKey[i], firstKeyConstant)
		intermediateByte = utils.CalculateXor(intermediateByte, secondKeyConstant)

		intermediateKey[i] = intermediateByte
	}

	return intermediateKey
}

// KeySchedule : Cyclic permutation
func (k key) KeySchedule(intermediateKey [][]byte) [][]byte {
	// b3 b15 b12 b9
	firstBytes := []byte{intermediateKey[0][2], intermediateKey[3][2], intermediateKey[2][3], intermediateKey[2][0]}
	// b4 b14 b13 b7
	secondBytes := []byte{intermediateKey[0][3], intermediateKey[3][1], intermediateKey[3][0], intermediateKey[1][2]}
	// b2 b16 b11 b5
	thirdBytes := []byte{intermediateKey[0][1], intermediateKey[3][3], intermediateKey[2][2], intermediateKey[1][0]}
	// b1 b10 b8 b6
	fourthBytes := []byte{intermediateKey[0][0], intermediateKey[2][1], intermediateKey[1][3], intermediateKey[1][1]}

	keySchedule := [][]byte{firstBytes, secondBytes, thirdBytes, fourthBytes}

	return keySchedule
}

// GenerateSBox : XOR key schedule with pi
func (k key) GenerateSBox(keySchedule [][]byte) [][]byte {
	// Initialize using fractional portion of pi (.1415926535897932384626433832795028)
	sBox := [][]byte{{141, 59, 26, 53}, {58, 97, 93, 238}, {46, 26, 43, 38}, {32, 79, 50, 28}}

	var newSBox [][]byte
	for i, b := range keySchedule {
		subArray := utils.CalculateXor(b, sBox[i])
		newSBox = append(newSBox, subArray)
	}

	return newSBox
}
