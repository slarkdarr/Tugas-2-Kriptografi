package internal

type (
	Cipher interface {
		Encrypt(plaintext string) []uint8
		Decrypt(block []uint8) string
	}
)

type (
	Key interface {
		Generate() string
	}
)

type (
	Substitution interface {
		Execute(block []uint8) []uint8
	}
)

type (
	Permutation interface {
		Execute(block []uint8) []uint8
	}
)
