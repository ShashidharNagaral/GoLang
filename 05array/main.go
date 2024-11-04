package main

import "fmt"

func main() {

	var groceryList [5]string // array declaration with default value ""

	groceryList[0] = "milk"
	groceryList[1] = "bread"
	groceryList[3] = "butter"

	fmt.Println("groceryList: ", groceryList)
	fmt.Println("groceryList size: ", len(groceryList))
	/*
		NOTE: by default,
		[]string has "" is added
		[]int 0 is added
	*/

	var fruitList = [3]string{"mango", "banana", "chickoo"} // declaration with initialization

	fmt.Println("fruit list:", fruitList)

	flowers := [...]string{"lily", "sunflower", "rose"} // length infered
	fmt.Println(flowers)

	/*
		NOTE: An array type definition specifies a length and an element type. For example, the type [4]int
		represents an array of four integers. An array’s size is fixed; its length is part of
		its type ([4]int and [5]int are distinct, incompatible types).
	*/

	// POINTER TO ARRAY
	arr := [3]int{1, 2, 3}

	// var p *int = arr // In C/C++, this is allowed since array is refernce

	var p *[3]int = &arr
	fmt.Println("*p", *p)
	fmt.Printf("address of arr: %p\n", &arr)
	fmt.Printf("address of arr[0]: %p\n", &arr[0])
	fmt.Printf("value of p: %p\n", p)
	fmt.Printf("address of p: %p\n", &p)

	// passing go array to a function
	A := [5]int{1, 4, 2, 1, 3}
	fmt.Println("value at index 3 before function call: ", A[3])
	updateArrayValueAtIndex(3, 100, A)
	fmt.Println("value at index 3 after function call: ", A[3])

	// NOTE: passing array to a function is call by value and not call by referece
	// i.e. entire new copy of array in passed to a function

	// to make it pass by reference we have to explicitly create a pointer of [n]T type
	fmt.Println("value at index 3 before function call: ", A[3])
	updateArrayAtPosition(3, 100, &A)
	fmt.Println("value at index 3 after function call: ", A[3])
}

func updateArrayValueAtIndex(index int, value int, arr [5]int) {
	arr[index] = value
}

func updateArrayAtPosition(index int, value int, p *[5]int) {
	(*p)[index] = value
}
