package internal

type (
	Cipher interface {
		Encrypt(plaintext, externalKey string) string
		Decrypt(ciphertext, externalKey string) string
	}

	Key interface {
		Generate() []uint32
	}

	Substitution interface {
		Execute(chunk uint32) uint32
	}

	Permutation interface {
		Execute(chunk uint32) uint32
	}
)
