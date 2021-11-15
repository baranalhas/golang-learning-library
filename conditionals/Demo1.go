package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900
	if cekilmekIstenen > hesap {
		fmt.Println("Para yetersiz")
	} else if hesap > cekilmekIstenen {
		hesap = hesap - cekilmekIstenen
		fmt.Println("İşlem başarılı Bakiye", hesap)
		fmt.Println("İşlem başarılı Bakiye ", fmt.Sprintf("%v", hesap))
	} else {
		fmt.Println("Bir hata oluştu")
	}
}
