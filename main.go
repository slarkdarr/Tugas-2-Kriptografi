package main

import (
	"fmt"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/impl"
)

func main() {
	c := impl.NewCipher("asdfgh")
	encrypted := c.Encrypt("hai mari berhimpun123")
	fmt.Println("Encrypted: ", encrypted)
	decrypted := c.Decrypt(encrypted)
	fmt.Println("Decrypted: ", decrypted)
}
