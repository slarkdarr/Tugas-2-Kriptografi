package internal

type (
	Cipher interface {
		Encrypt(plaintext, externalKey string) string
		Decrypt(ciphertext, externalKey string) string
	}

	Key interface {
		Generate() [][]byte
	}

	Substitution interface {
		Execute(chunk []byte) []byte
	}

	Permutation interface {
		Execute(chunk []byte) []byte
	}
)
