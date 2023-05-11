package main

import (
	"fmt"

	mymath "example.com/test/math"
	mytry "example.com/test/try"
)

func main() {
	fmt.Println("Hi")
	fmt.Println(mymath.Add(2, 2))
	// mytry.Pointer()
	mytry.Print("Bob", 5, "brown")
	mytry.Print("Blacki", 5, "black")
}
