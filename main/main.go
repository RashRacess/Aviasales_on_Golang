package main

import (
	"aviasales/main/authentification"
	"aviasales/main/interfaces"
	parsingtxt "aviasales/main/parsing_txt"
	"bufio"
	"fmt"
	"os"

	fl "github.com/RashRacess/Flight"
	p "github.com/RashRacess/Person"
)

func main() {
	_ = fl.Flight{}
	_ = p.Person{}

	scanner := bufio.NewScanner(os.Stdin)
	var surname string
	var name string

	fmt.Print("Enter surname: ")
	_ = scanner.Scan()
	surname = scanner.Text()

	fmt.Print("Enter name : ")
	_ = scanner.Scan()
	name = scanner.Text()

	persons, _ := parsingtxt.Parsing_txt_persons("./parsing_txt/persons.txt")

	person, err := authentification.Entering(p.CreatePerson(0, surname, name, "user", 0), &persons)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if person.GetRole() == "admin" {
		ai := interfaces.CreateAdminInterface(person)
		ai.Admin()
	} else {
		ui := interfaces.CreateUserInterface(person)
		ui.User()
	}
}
