package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 User = User{"Shashidhar", 26}
	fmt.Println(user1)
	fmt.Printf("%+v \n", user1)
	fmt.Printf("Name: %v \n", user1.Name)
}
