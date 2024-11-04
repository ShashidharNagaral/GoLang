package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(6)
	fmt.Println(r)
	switch r {
	case 1:
		fmt.Printf("this is 1")
		fallthrough // this all pass below case
	case 2:
		fmt.Printf("this is 2")
	case 3:
		fmt.Printf("this is 3")
	case 4:
		fmt.Printf("this is 4")
	case 5:
		fmt.Printf("this is 5")
	case 6:
		fmt.Printf("this is 6")
	default:
		fmt.Printf("invalid dice")
	}
}
