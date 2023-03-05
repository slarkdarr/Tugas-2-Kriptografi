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
	c := impl.NewCipher("t6w9z$C&F)H@McQf")
	testcase := []string{
		"kriptografi",
		"kr1ptografi",
		"christo daffa abc",
		"christo daffa abd",
		"christo viel daf",
		"christo viel d4f",
		"christo vieldaff",
		"khristo vieldaff",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
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
