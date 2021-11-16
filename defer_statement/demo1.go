package defer_statement

import "fmt"

func B() {
	defer C()
	defer A()
	// son gelen ilk gider mantığıyla ilk çalışacak olan defer A'dır
	fmt.Println("b fonksiyonu çalıştı")
}

func C() {
	fmt.Println("c fonksiyonu çalıştı")
}

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}
