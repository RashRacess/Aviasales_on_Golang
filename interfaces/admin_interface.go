package interfaces

import (
	"aviasales/essence"
	parsinginformation "aviasales/parsing_information"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Admin struct {
	essence.Person
}

func AdminInterface(person essence.Person) {
	admin := Admin{person}
	var flights = parsinginformation.Parse_txt_flights("../parsing_information/flights.txt")
	var users = parsinginformation.Parse_txt_persons("../parsing_information/persons.txt")
	scanner := bufio.NewScanner(os.Stdin)
	var isWork bool = true

	for isWork {
		fmt.Println("1. Show profile")
		fmt.Println("2. Show users")
		fmt.Println("3. Delete user")
		fmt.Println("4. Add user")

		fmt.Println("5. Show flights")
		fmt.Println("6. Delete flight")
		fmt.Println("7. Add flight")
		fmt.Println("8. Exit")

		_ = scanner.Scan()
		choise, err := strconv.Atoi(scanner.Text())
		if err != nil {
			choise = 0
		}
		switch choise {
		case 1:
			fmt.Printf("%s - %s - %s", admin.GetName(), admin.GetSurname(), admin.GetRole())
			_ = scanner.Scan()

		case 2:
			for _, v := range users {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 5:
			for _, v := range flights {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 8:
			fmt.Println("Goodbye")
			isWork = false
		default:
		}
	}
}
