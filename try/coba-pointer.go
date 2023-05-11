package try

import (
	"fmt"
)

func Pointer() {

	i := 42

	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println("++++++++++++++++")

	p := &i
	*p = 20
	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println("_____________")

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	fmt.Println("++++++++++++++++")
	z := p
	*z = 30

	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println("_____________")

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	fmt.Println("_____________")

	fmt.Println(z)
	fmt.Println(*z)
	fmt.Println(&z)
}
