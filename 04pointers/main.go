package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var govar int
	var govar_p *int

	fmt.Println(govar, govar_p)

	govar = 12
	govar_p = &govar

	fmt.Println("govar value: ", govar, " and its address: ", govar_p)
	fmt.Println("we can access the value at variable using pointer as, *govar_p", *govar_p)

	// update the value using pointer variable
	*govar_p = 21
	fmt.Println(govar)

	var govar_p_p **int
	govar_p_p = &govar_p // double pointer which stores address of another pointer

	fmt.Println("govar", govar)
	fmt.Println("govar_p", *govar_p)
	fmt.Println("govar_p_p", **govar_p_p)
}
