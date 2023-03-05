package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Permutation(t *testing.T) {
	t.Run("Permute Forward", func(tt *testing.T) {
		p := NewPermutation()
		input := []byte{0x12, 0x34, 0x56, 0x78}
		output := []byte{0x1d, 0xba, 0x13, 0x80}

		outputResult := p.Execute(input)
		assert.EqualValues(tt, output, outputResult)
	})
	//t.Run("Permute Reverse", func(tt *testing.T) {
	//	p := NewPermutation()
	//	input := []byte{0x12, 0x34, 0x56, 0x78}
	//	output := []byte{0x1d, 0xba, 0x13, 0x80}
	//	inputResult := p.Reverse(output)
	//	assert.EqualValues(tt, input, inputResult)
	//})
	//t.Run("Permute Forward-Reverse Random", func(tt *testing.T) {
	//	p := NewPermutation()
	//	for x := 0; x < 50; x++ {
	//		num := rand.Uint32()
	//		data := []byte{0, 0, 0, 0}
	//		binary.BigEndian.PutUint32(data, num)
	//		output := p.Execute(data, true)
	//		input := p.Execute(output, false)
	//		resultNum := binary.BigEndian.Uint32(input)
	//		assert.EqualValues(tt, num, resultNum)
	//	}
	//})
}
