package main

import (
	"aviasales/essence"
	"aviasales/interfaces"
)

func main() {
	p := essence.NewPerson(1, "Kozyrev", "Dima", "user")
	interfaces.AdminInterface(p)
}
