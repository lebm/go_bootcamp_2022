package main

import "fmt"

func main() {
	// type comes after the variable name
	var x int = 7

	// same declaration as above with type inferred from the RHS value
	// Used when you do not know the initial value. Assume the default value of the type.
	var x2 = 7

	// short declaration
	// go can infer the type
	// works only in blocks code
	// Note that ":=" is a declaration, "=" is an atribution
	// Used when you know the initial value
	y := 10
	y = 9

	// variables must be used, or the compiler will complain
	fmt.Println(x, x2, y)

	// blank ientifier "_" can be used to avoid errors for not used variables
	// use with care
	var not_used = 100
	_ = not_used

	// Multiple declarations
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// Its an errot
	// car, cost := "Audi", 60000
	// Its ok, year does not exist yet
	car, year := "Volvo", 2018
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	// Uses "zero" values when not declared
	fmt.Println(salary, firstName, gender)

	// variables is allways initialized, if a RHS expression is not provided, go uses the "zero value" of the type
	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	_, _ = i, j

	// swap variables
	i, j = j, i
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

}
