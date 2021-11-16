package string_functions

//alias
import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Baran"
	fmt.Println(s.Count(isim, "a"))
	fmt.Println(s.Contains(isim, "a"))
	fmt.Println(s.Index(isim, "a"))
	fmt.Println(s.ToLower(isim), s.ToUpper(isim))
	fmt.Println(s.HasPrefix(isim, "Ba"), "<-->", s.HasSuffix(isim, "an"))
	harfler := []string{"b", "a", "r", "a", "n"}
	fmt.Println(s.Join(harfler, "."))
	fmt.Println(s.Replace(isim, "a", "", -1))
	fmt.Println(s.Replace(isim, "a", "", 1))
	fmt.Println(s.Split(isim, "a"))
	fmt.Println(s.Repeat(isim, 5))
}
