package main

// import other packages
import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, World!!!!!!!!!")

	// varuables and constants
	// string int float32 bool

	//var firstName string = "Jouni"

	//:= shorthand variables

	//lastName := "kiviperä"
	//age := 44

	// create but dont't assing

	var fullName string
	var age int
	var price float32
	var tf bool

	fmt.Println(fullName, age, price, tf)

	fullName = "Jouni Kiviperä"
	age = 41
	price = 20.00
	tf = true

	fmt.Println(fullName, age, price, tf)

	// multiple variable declaration

	var name1, name2 string = "Pekka", "Reima"
	fmt.Println(name1, name2)

	// constanst

	const peli string = "super mario"
	fmt.Println(peli)

	// multi const

	const (
		peli1 = "super mario"
		peli2 = "zelda"
		num   = "100"
	)

	fmt.Println(peli2)

	//fmt.Println(firstName)

}
