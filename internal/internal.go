package internal

type (
	Cipher interface {
		Encrypt(plaintext string) string
		Decrypt(ciphertext string) string
	}

	Key interface {
		Generate() [][]byte
	}

	Executor interface {
		Execute(chunk []byte) []byte
	}

	GroupExecutor interface {
		Executor
		ResetCount()
	}
)
