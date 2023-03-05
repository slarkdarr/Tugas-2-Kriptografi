package main

import (
	"encoding/base64"
	"fmt"
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal/impl"
	"strconv"
	"time"
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
		"Pak Rinaldi adalah seorang guru kelas kriptografi yang sangat terkenal di kota itu. Setiap minggu, ia selalu memberikan ceramah tentang berbagai teknik kriptografi yang digunakan untuk mengamankan informasi rahasia. Murid-muridnya sangat senang belajar dengan Pak Rinaldi karena ia selalu membuat pelajarannya menarik dan mudah dipahami.\n\nSuatu hari, saat pelajaran berlangsung, tiba-tiba terdengar suara gemerincing dari luar kelas. Pak Rinaldi keluar untuk memeriksanya dan ternyata sebuah kotak pos besar telah tiba. Ia membuka kotak tersebut dan menemukan beberapa mesin enkripsi yang sangat canggih di dalamnya.\n\nPak Rinaldi sangat senang dengan mesin-mesin tersebut dan mengajak murid-muridnya untuk mempelajarinya bersama-sama. Mereka belajar tentang bagaimana cara kerja mesin dan cara menggunakannya untuk mengamankan pesan-pesan rahasia. Pak Rinaldi sangat bangga melihat betapa antusiasnya murid-muridnya dalam mempelajari kriptografi.\n\nHari berikutnya, Pak Rinaldi memberikan tugas kepada murid-muridnya untuk mencoba menggunakan mesin enkripsi yang baru mereka pelajari untuk mengamankan pesan rahasia. Murid-muridnya sangat semangat dan berusaha keras untuk menyelesaikan tugas tersebut dengan baik.\n\nSetelah beberapa hari, Pak Rinaldi memeriksa hasil tugas murid-muridnya dan sangat senang dengan hasilnya. Mereka berhasil mengamankan pesan rahasia dengan menggunakan mesin enkripsi yang baru mereka pelajari.\n\nPak Rinaldi sangat bangga dengan murid-muridnya dan senang melihat betapa mereka begitu antusias belajar tentang kriptografi. Ia berharap suatu hari nanti mereka dapat menjadi ahli kriptografi yang handal dan membantu mengamankan informasi rahasia di seluruh dunia.",
	}

	// Incase mau coba string gede-gede, dari file aja, tinggal uncomment
	//f, _ := os.Open("./Termius.dmg")
	//b, _ := io.ReadAll(f)
	//testcase = append(testcase, string(b))

	for i, each := range testcase {
		start := time.Now()
		enc := c.Encrypt(each)
		encT := time.Now().Sub(start)
		start = time.Now()
		dec := c.Decrypt(enc)
		decT := time.Now().Sub(start)
		if len(each) < 200 {
			fmt.Println("Testcase", i, "->", each)
			fmt.Println("E\t: ", enc)
			fmt.Println("E (HEX)\t: ", getArrayOfHex(enc))
			fmt.Println("D\t: ", dec)
		} else {
			fmt.Println("Testcase", i, "("+strconv.Itoa(len(each))+" characters)")
		}
		fmt.Println("E(T)\t: ", encT)
		fmt.Println("D(T)\t: ", decT)
	}
}
