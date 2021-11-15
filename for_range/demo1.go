package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "Kocaeli"}

	for i, sehir := range sehirler {
		fmt.Println(i, " ", sehir)
	}

}
