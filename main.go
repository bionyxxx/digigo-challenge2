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
			str := "САШАРВО"
			bytePosition := 0

			for i, char := range str {
				//mencari posisi byte pada range string
				bytePosition = len([]byte(str[:i]))
				//cek untuk nilai byte yang genap
				if bytePosition%2 == 0 {
					fmt.Printf("character U+%04X '%s'  starts at byte position %d\n", char, string(char), bytePosition)
				}
			}

		default:
			println("Nilai j adalah", j)
		}
	}
}
