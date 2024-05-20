package user

import "fmt"

type User interface {
	login()
	logout()
}

type Admin struct {
	Name string
}

func (a *Admin) login() {
	fmt.Println("Admin login")
}

func (a *Admin) logout() {
	fmt.Println("Admin logout")
}

type Customer struct {
	Name string
}

func (c *Customer) login() {
	fmt.Println("Customer login")
}

func (c *Customer) logout() {
	fmt.Println("Customer logout")
}

func Process_order(user User) {
	user.login()
	fmt.Println("Processing order...")
	user.logout()
}
