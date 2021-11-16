package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir değer giriniz !")
	}
	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı giriniz", nil
	}
	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı giriniz", nil
	}
	return "Bildiniz", nil
}

func Demo2() {
	msg, _ := TahminEt(50)
	fmt.Println(msg)
	fmt.Println(TahminEt(10))
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-10))
}
