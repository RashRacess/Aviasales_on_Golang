package main

import (
	"aviasales/main/authentification"
	parsingtxt "aviasales/main/parsing_txt"
	"fmt"

	fl "github.com/RashRacess/Flight"
	p "github.com/RashRacess/Person"
)

func main() {
	_ = fl.Flight{}
	_ = p.Person{}

	var surname string
	var name string

	fmt.Print("Enter surname: ")
	fmt.Scanln(&surname)
	fmt.Print("Enter name : ")
	fmt.Scanln(&name)

	flights, _ := parsingtxt.Parsing_txt_flights("../flights.txt")
	_ = flights

	persons, _ := parsingtxt.Parsing_txt_persons("../persons.txt")
	_ = persons

	person, err := authentification.Entering(p.CreatePerson(0, surname, name, "user", 0), &persons)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(person)
	}
}
