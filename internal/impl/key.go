package impl

import "github.com/slarkdarr/Tugas-2-Kriptografi/internal"

type (
	key struct {
		externalKey string
	}
)

func NewKey(externalKey string) internal.Key {
	return &key{externalKey}
}

func (k key) Generate() string {
	//TODO implement me
	panic("implement me")
}
