package maps

import "fmt"

func Demo1(value string) string {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	delete(sozluk, "book")

	deger, varMi := sozluk["table"]
	fmt.Println(deger, varMi)
	deger, varMi = sozluk["book"]
	fmt.Println(deger, varMi)

	sozluk2 := map[string]string{"glass": "bardak", "phone": "telefon"}

	fmt.Println(sozluk2)
	return sozluk[value]
}
