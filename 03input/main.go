package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter rating from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	rating, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rating + 1)
	}
}
