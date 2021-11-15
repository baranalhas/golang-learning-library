package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	isimler[0] = "Baran"
	isimler[1] = "Engin"
	isimler[2] = "Salih"

	fmt.Println(isimler)
	isimler = append(isimler, "Derin")
	fmt.Println(isimler)
}
