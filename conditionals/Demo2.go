package conditionals

import "fmt"

func Demo2() {

	var sayi1, sayi2, sayi3 int = 10, 5, 18

	var theBiggest int = sayi1
	if theBiggest < sayi2 {
		theBiggest = sayi2
	}
	if theBiggest < sayi3 {
		theBiggest = sayi3
	}
	fmt.Println(theBiggest)
}
