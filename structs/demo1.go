package structs

import "fmt"

func Demo1() {
	fmt.Println(product{name: "Bilgisayar", unitPrice: 5000, brand: "Monster"})
	fmt.Println(product{"Bilgisayar", 5000, "Monster", 0})
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
