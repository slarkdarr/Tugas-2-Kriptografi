package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Substitution(t *testing.T) {
	//t.Run("Create Sbox", func(tt *testing.T) {
	//	var sbox [256]byte
	//	for i := 0; i < 256; i++ {
	//		sbox[i] = byte(i)
	//	}
	//
	//	for k := 0; k < 16; k++ {
	//		for i := 255; i > 0; i-- {
	//			j := rand.Intn(i + 1)
	//			sbox[i], sbox[j] = sbox[j], sbox[i]
	//		}
	//	}
	//	fmt.Printf("%#v", sbox)
	//})
	t.Run("Substitute Forward", func(tt *testing.T) {
		p := NewSubstitution()
		input := []byte{0x12, 0x34, 0x56, 0x78}
		output := []byte{0x7f, 0x1f, 0x2a, 0x33}

		outputResult := p.Execute(input)
		assert.EqualValues(tt, output, outputResult)
	})
	//t.Run("Substitute Reverse", func(tt *testing.T) {
	//	p := NewSubstitution()
	//	input := []byte{0x12, 0x34, 0x56, 0x78}
	//	output := []byte{0x7f, 0x1f, 0x2a, 0x33}
	//	inputResult := p.Execute(output, false)
	//	assert.EqualValues(tt, input, inputResult)
	//})
	//t.Run("Substitute Forward-Reverse Random", func(tt *testing.T) {
	//	p := NewSubstitution()
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
