package variables

import "fmt"

//export edilen paketler büyük harfle başlar
//camelCase PascalCase
func Demo1() {
	var text string = "Hello Turkey!"

	fmt.Println(text)

	fmt.Print("deneme")
	fmt.Println("TEST")

	var int_test int = 10
	fmt.Print("test ", int_test, " asdfasdf ")
	var integ int
	integ = 5
	fmt.Println(integ)

	var float_test float32 = 0.1
	fmt.Println(float_test)

	testvar := 25
	fmt.Println(testvar)
	fmt.Printf("veri tipi : %T \nTest\n", testvar)

	var status bool = false
	metin1 := "engin"
	metin2 := "ahmet"
	status = metin1 == metin2
	if status {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
