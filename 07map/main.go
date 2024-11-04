package main

import "fmt"

// var <variable_name> map[<key_type>]<value_type>
var m map[string]int

func main() {
	// declare a map

	m = make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println(m)

	for i, v := range m {
		fmt.Printf("key: %v value: %d\n", i, v)
	}
}
