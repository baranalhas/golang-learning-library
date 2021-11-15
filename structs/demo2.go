package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi: ", c.firstName)
}
func (c customer) update() {
	fmt.Println("Güncellendi: ", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Baran", lastName: "A", age: 23}
	c.save()
	c.update()
}
