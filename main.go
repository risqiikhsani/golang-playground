package main

import (
	"fmt"

	mymath "example.com/test/math"
	mytry "example.com/test/try"
)

type Car struct {
	name       string
	publisher  string
	car_type   string
	drive_type string
}

func main() {
	fmt.Println("Hi")
	fmt.Println(mymath.Add(2, 2))
	// mytry.Pointer()
	mytry.Print("Bob", 5, "brown")
	mytry.Print("Blacki", 5, "black")

	var a [5]Car
	a[0] = Car{"M4", "BMW", "sedan", "manual"}
	a[1] = Car{"Civic", "Honda", "Sedan", "manual"}
	fmt.Println(a)
}
