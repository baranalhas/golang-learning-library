package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "Kocaeli"
	sehirler[3] = "İzmir"
	sehirler[4] = "Adana"
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
}
