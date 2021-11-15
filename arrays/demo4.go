package arrays

import "fmt"

func Demo4() {
	var sayilar [2][3]int
	sayilar[0][0] = 1
	sayilar[0][1] = 2
	sayilar[0][2] = 3

	sayilar[1][0] = 4
	sayilar[1][1] = 5
	sayilar[1][2] = 6

	//fmt.Print(sayilar)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j], " ")
		}
		fmt.Print("\n")
	}
}
