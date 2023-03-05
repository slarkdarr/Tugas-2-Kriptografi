package utils

import (
	"encoding/binary"
	"math"
)

func CalculateXor(a, b []byte) []byte {
	x1 := binary.BigEndian.Uint32(a)
	x2 := binary.BigEndian.Uint32(b)

	xor := x1 ^ x2

	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, xor)

	return result
}

func CalculateAddMod32(a, b []byte) []byte {
	x1 := binary.BigEndian.Uint32(a)
	x2 := binary.BigEndian.Uint32(b)

	add := (x1 + x2) % (math.MaxUint32)

	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, add)

	return result
}
