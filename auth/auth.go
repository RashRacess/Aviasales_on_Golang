package auth

import (
	"aviasales/essence"
	"aviasales/interfaces"
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Froggy() {
	var surname, name, choise string
	var filename string = "../parsing_information/persons.txt"
	var count_of_entered int

	for count_of_entered != 2 {
		fmt.Println("Enter surname and name")
		count_of_entered, _ = fmt.Scanln(&surname, &name)
	}

	fmt.Println("Hello", name)
	fmt.Println("If you want to sign up or sign up")
	for count_of_entered != 1 {
		fmt.Println("up/in")
		count_of_entered, _ = fmt.Scanln(&choise)
		if choise != "up" && choise != "in" {
			count_of_entered = 0
		}
	}

	if choise == "up" {
		p, err := SignUp(surname, name, filename)
		if err != nil {
			log.Fatalln(err)
		}

		if p.GetRole() == "user" {
			interfaces.UserInterface(p)
		} else if p.GetRole() == "admin" {
			interfaces.AdminInterface(p)
		} else {
			log.Println("role:", p.String())
		}

	} else {
		p, err := SignIn(surname, name, filename)
		if err != nil {
			log.Fatalln(err)
		}

		if p.GetRole() == "user" {
			interfaces.UserInterface(p)
		} else if p.GetRole() == "admin" {
			interfaces.AdminInterface(p)
		} else {
			log.Println("role:", p.String())
		}
	}
}

func SignIn(surname, name, filename string) (essence.Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return essence.Person{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ";")
		if words[1] == surname && words[2] == name {
			id, err := strconv.Atoi(words[0])
			if err != nil {
				id = 0
			}
			return essence.NewPerson(id, words[1], words[2], words[3]), nil
		}
	}
	return essence.Person{}, errors.New("user is not exist")
}

func SignUp(surname, name, filename string) (essence.Person, error) {
	file, err := os.OpenFile(filename, os.O_APPEND, 0644)
	if err != nil {
		return essence.Person{}, err
	}
	defer file.Close()

	p, err := SignIn(surname, name, filename)
	if err != nil {
		p = essence.NewPerson(rand.Intn(10000), surname, name, "user")
		file.WriteString(p.StringWithSeparator(";") + "\n")
		return p, nil
	}
	return p, errors.New("user is exists")
}
