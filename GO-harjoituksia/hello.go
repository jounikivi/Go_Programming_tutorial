package main

// import variables
import (
	"fmt"
)

func main() {

	// variables and Constants
	// strings, int, float32, bool

	//var firstName string = "Jouni"

	// := shorthand variables

	//lastName := "Kiviperä"
	//age := 41

	// Create but dont's assign

	var fullName string
	var age int
	var price float32
	var tf bool

	fmt.Println(fullName, age, price, tf)

	fullName = "Jouni Kiviperä"
	age = 41
	price = 19.99
	tf = true

	fmt.Println(fullName, age, price, tf)

}
