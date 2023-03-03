package impl

import "github.com/slarkdarr/Tugas-2-Kriptografi/internal"

type (
	cipher struct {
		externalKey string
	}
)

func NewCipher(externalKey string) internal.Cipher {
	return &cipher{externalKey: externalKey}
}

func (c cipher) Encrypt(plaintext string) []uint8 {
	//TODO implement me
	panic("implement me")
}

func (c cipher) Decrypt(block []uint8) string {
	//TODO implement me
	panic("implement me")
}
