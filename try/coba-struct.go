package try

import (
	"fmt"
)

type Kucing struct {
	name  string
	age   int
	color string
}

func Print(name string, age int, color string) {
	test := Kucing{name, age, color}
	fmt.Println(test.name)
	fmt.Println(test.age)
	fmt.Println(test.color)
}
