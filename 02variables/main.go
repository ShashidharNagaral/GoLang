package main

import (
	"fmt"
)

// globalVar := "no var declaration is not allowed in global scope"

var globalToken = "global token: package level scope (global scope)" // this is package level variable declaration

func main() {
	// var govar int = 12 // declaring a variable without using it will throw an error

	// int
	var intVar int = 12
	fmt.Println(intVar)
	fmt.Printf("Type: %T \n", intVar)

	// string
	var stringVar string = "this is a string"
	fmt.Println(stringVar)
	fmt.Printf("Type: %T \n", stringVar)

	// float
	var floatVar float32 = 22.22
	fmt.Println(floatVar)
	fmt.Printf("Type: %T \n", floatVar)

	// default values
	var defaultInt int // 0
	fmt.Println(defaultInt)

	var defaultString string // ""
	fmt.Println(defaultString)

	var defaultFloat32 float32 // 0
	fmt.Println(defaultFloat32)

	// implicit type
	var Istring = "this is a string type" // lexer will declare variable1 as string type
	fmt.Println(Istring)
	fmt.Printf("Type: %T \n", Istring)

	var Iint = 12 // int
	fmt.Println(Iint)
	fmt.Printf("Type: %T \n", Iint)

	var Ifloat = 12.2 // float64
	fmt.Println(Ifloat)
	fmt.Printf("Type: %T \n", Ifloat)

	// no var declaration
	// NOTE: this is only allowed inside method
	count := 12
	fmt.Println(count)
	fmt.Printf("Type: %T \n", count)

	fmt.Println(globalToken)

	// function scope (local scope)

	localVariable := "this is function scope variable"
	fmt.Println(localVariable)

	goFunction()

	// block scope
	if true {
		blockVar := "this is block scope variable"
		fmt.Println(blockVar)
	}
	// fmt.Println(blockVar) // blockVar is not accessible outside its declared block

}

func goFunction() {
	fmt.Println("localVariable from main function cannot be accessed in goFunction")
}
