package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	fmt.Println("")
	/*project.GetAllProducts()
	project.AddProduct()
	project.GetAllProducts()*/

	project.AddProduct()

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

	//restful.Demo2()
	//string_functions.Demo1()
	//fmt.Println(error_handling.TahminEt2(102))
	//interfaces.Demo3()
	//defer_statement.B()
	//defer_statement.Test()
	//defer_statement.Demo3()

	//interfaces.Demo2()
	//
	/* thread
	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)
	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn


	fmt.Println("Çarpım: ", ciftSayiToplam*tekSayiToplam)

	*/
	/*
			//structs.Demo1()
		structs.Demo2()
			sayilar := []int{1, 2, 3}
			pointers.Demo2(sayilar)
			fmt.Println("Maindeki sayı ", sayilar[0])*/

	/* pointers

	sayi := 20
	pointers.Demo1(&sayi)
	fmt.Println("Maindeki sayi: ", sayi)*/

	//variables.Demo1()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//loops.FriendNumbers()
	/*
		var number int
		fmt.Print("Bir sayı gir: ")
		fmt.Scanln(&number)
		fmt.Print("Girilen sayı: ", number)*/

	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//slices.Demo1()
	//slices.Demo2()
	/*
			sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

			fmt.Println(sonuc1)
			fmt.Println(sonuc2)
			fmt.Println(sonuc3)
			//fmt.Println(sonuc4)

			fmt.Println(functions.ToplaVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9))

			sayilar := []int{1, 4, 2, 6, 7, 8, 9, 1, 2, 3}
			fmt.Println(functions.ToplaVariadic(sayilar...))

		fmt.Println(maps.Demo1("book"))
		for_range.Demo1()
		for_range.Demo2()
		for_range.Demo3()*/

}
