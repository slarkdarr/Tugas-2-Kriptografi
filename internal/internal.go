package internal

type (
	Cipher interface {
		Encrypt(plaintext string) string
		Decrypt(ciphertext string) string
	}

	Key interface {
		Generate() [][]byte
	}

	Substitution interface {
		Execute(chunk []byte, encrypt bool) []byte
	}

	Permutation interface {
		Execute(chunk []byte, encrypt bool) []byte
	}
)
