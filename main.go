package main

import (
	"fmt"

	us "github.com/juandiegomoya/kata/user"
)

func main() {
	customer := us.Customer{Name: "Juan"}
	admin := us.Admin{Name: "Alice"}

	fmt.Println("Processing order for Customer:")
	us.Process_order(&customer)

	fmt.Println("\nProcessing order for Admin:")
	us.Process_order(&admin)
}
