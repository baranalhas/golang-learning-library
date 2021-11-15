package loops

import "fmt"

func FriendNumbers() {
	sayi1 := 220
	sayi2 := 284
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 += i
		}
	}

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Print(sayi1, " ve ", sayi2, " arkadaş sayılardır")
	} else {
		fmt.Print(sayi1, " ve ", sayi2, " arkadaş sayılardır")
	}
}
