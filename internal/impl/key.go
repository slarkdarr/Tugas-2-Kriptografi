package impl

import (
	"encoding/binary"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/utils"
	"math"
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
	subKeys := k.GenerateSubKeys(keySchedule, sBox)

	return subKeys
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

// GenerateSubKeys : Generate 32 subkeys (4 bytes each)
func (k key) GenerateSubKeys(keySchedule, sBox [][]byte) [][]byte {
	var subKeys [][]byte
	var subKey []byte

	for i := 0; i < 32; i++ {
		// auto mod 256
		if i < 4 {
			subKey = append(subKey, keySchedule[i][0]+sBox[i][0])
			subKey = append(subKey, keySchedule[i][1]+sBox[i][1])
			subKey = append(subKey, keySchedule[i][2]+sBox[i][2])
			subKey = append(subKey, keySchedule[i][3]+sBox[i][3])
		} else if i >= 4 && i < 8 {
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][1])
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][2])
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][3])
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][0])
		} else if i >= 8 && i < 12 {
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][2])
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][3])
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][0])
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][1])
		} else if i >= 12 && i < 16 {
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][3])
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][0])
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][1])
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][2])
		} else if i >= 16 && i < 20 {
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][0])
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][1])
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][2])
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][3])
		} else if i >= 20 && i < 24 {
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][0])
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][1])
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][2])
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][3])
		} else if i >= 24 && i < 28 {
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][0])
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][1])
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][2])
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][3])
		} else if i >= 28 && i < 32 {
			subKey = append(subKey, keySchedule[i%4][3]+sBox[i%4][3])
			subKey = append(subKey, keySchedule[i%4][2]+sBox[i%4][2])
			subKey = append(subKey, keySchedule[i%4][1]+sBox[i%4][1])
			subKey = append(subKey, keySchedule[i%4][0]+sBox[i%4][0])
		}

		subKeys = append(subKeys, subKey)
		subKey = nil
	}

	var finalSubKeys [][]byte
	for i, subKey := range subKeys {
		finalSubKeyFloat := math.Sqrt(math.Pi / float64(i))
		finalSubKey := k.Float64Bytes(finalSubKeyFloat)
		finalSubKey = utils.CalculateXor(subKey, finalSubKey)

		finalSubKeys = append(finalSubKeys, finalSubKey)
	}

	return finalSubKeys
}

func (k key) Float64Bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
