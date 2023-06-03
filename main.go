package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// mymath "github.com/risqiikhsani/golang-playground/math"
// mytry "github.com/risqiikhsani/golang-playground/try"

type Car struct {
	name       string
	publisher  string
	car_type   string
	drive_type string
}

type Cat struct {
	gorm.Model
	name string
	age  int
}

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to open database")
	}

	db.AutoMigrate(&Cat{})
	a := Cat{name: "bob", age: 1}
	b := Cat{name: "whiti", age: 2}
	c := Cat{name: "blacky", age: 1}
	d := []Cat{{name: "soy", age: 1}, {name: "mey", age: 1}, {name: "mu", age: 1}}

	db.Create(&a)
	db.Create(&b)
	db.Create(&c)
	db.Create(&d)

	var test Cat
	db.First(&test, 1)

	// fmt.Println("Hi")
	// fmt.Println(mymath.Add(2, 2))
	// // mytry.Pointer()
	// mytry.Print("Bob", 5, "brown")
	// mytry.Print("Blacki", 5, "black")

	// var a [5]Car // this is array
	// a[0] = Car{"M4", "BMW", "sedan", "manual"}
	// a[1] = Car{"Civic", "Honda", "Sedan", "manual"}
	// fmt.Println(a)

	// var b []Car // this is slice
	// var1 := Car{"M4 2020", "BMW", "sedan", "manual"}
	// var2 := Car{"Civic 1999", "Honda", "Sedan", "manual"}
	// b = append(b, var1, var2, Car{"Civic 2019", "Honda", "Sedan", "manual"})
	// fmt.Println(b)

	// a := 1
	// for b := 0; b < 5; b++ {
	// 	a++
	// 	fmt.Println(a)
	// }

	// istrue := false
	// if v := mymath.Add(2, 2); v < 5 {
	// 	istrue = true
	// }
	// fmt.Println(istrue)

	defer fmt.Println("test complete")
	// fmt.Println("test")

	// a, b := 1, 2
	// p := &a
	// fmt.Println(&a)
	// fmt.Println(*p)
	// fmt.Println(b)
	// fmt.Println(&p)
	// *p = *p + 2
	// fmt.Println(a)

}
