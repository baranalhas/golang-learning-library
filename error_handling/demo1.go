package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("./error_handling/demo1.txt")
	//dosya var err = nill
	if err != nil {
		fmt.Println("Dosya bulunamadı")
	} else {
		fmt.Println(f.Name())
	}
}
