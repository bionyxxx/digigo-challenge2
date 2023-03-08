package main

import "fmt"

func main() {
	//DUTTA FACHREZY

	//loop variable i
	for i := 0; i < 5; i++ {
		println("Nilai i adalah", i)
	}

	//loop variable j
	for j := 0; j <= 10; j++ {
		switch j {
		case 5:
			//Cara pengerjaan hanya apa yang ada di pikiran saya.
			str := "C A Ш A P B O"
			bytePosition := 0

			for i, char := range str {
				//mencari posisi byte pada range string
				bytePosition = len([]byte(str[:i]))
				//cek untuk nilai byte yang genap
				if bytePosition%2 == 0 {
					fmt.Printf("character U+%04X '%s'  starts at byte position %d\n", char, string(char), bytePosition)
				}
			}

			//Notes problem
			//output yang dihasilkan pada saat nilai J == 5 adalah:
			//character U+0043 'C'  starts at byte position 0
			//character U+0041 'A'  starts at byte position 2
			//character U+0428 'Ш'  starts at byte position 4
			//character U+0020 ' '  starts at byte position 6
			//character U+0020 ' '  starts at byte position 8
			//character U+0020 ' '  starts at byte position 10
			//character U+0020 ' '  starts at byte position 12

			//Tidak tahu kenapa setelah char Ш (Cyrillic) selalu terjadi kesalahan output pada char yang tersisa,
			//tapi saya sudah berusahan sebisanya, mungkin jika ada solusi bisa dibahas di sesi selanjutnya

		default:
			println("Nilai j adalah", j)
		}
	}
}
