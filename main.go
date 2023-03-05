package main

import (
	"encoding/base64"
	"fmt"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/impl"
)

func getArrayOfHex(text string) []string {
	data, _ := base64.StdEncoding.DecodeString(text)
	hexArr := make([]string, len(data))
	for i, b := range data {
		hexArr[i] = fmt.Sprintf("%02X", b)
	}
	return hexArr
}

func main() {
	c := impl.NewCipher("asdfgh")
	testcase := []string{
		"kriptografi",
		"kr1ptografi",
		"christo daffa abc",
		"christo daffa abd",
		"christo viel daf",
		"christo viel d4f",
		"christo vieldaff",
		"khristo vieldaff",
	}
	for i, each := range testcase {
		fmt.Println("Testcase", i, "->", each)
		enc := c.Encrypt(each)
		fmt.Println("E\t: ", enc)
		fmt.Println("E (HEX)\t: ", getArrayOfHex(enc))
		dec := c.Decrypt(enc)
		fmt.Println("D\t: ", dec)
	}
}
