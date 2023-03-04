package utils

import "math"

func CalculateXor(a, b uint32) uint32 {
	return a ^ b
}

func CalculateAddMod32(a, b uint32) uint32 {
	return (a + b) % (uint32(math.MaxUint32) + 1)
}
