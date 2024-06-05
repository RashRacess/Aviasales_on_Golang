package interfaces

import (
	parsingtxt "aviasales/main/parsing_txt"
	"bufio"
	"fmt"
	"os"
	"strconv"

	f "github.com/RashRacess/Flight"
	p "github.com/RashRacess/Person"
)

type AdminInterface struct {
	admin             p.Person
	flights           []f.Flight
	persons           []p.Person
	current_person_id int
	current_flight_id int
}

func CreateAdminInterface(p p.Person) AdminInterface {
	ai := AdminInterface{}
	ai.admin = p
	flights, _ := parsingtxt.Parsing_txt_flights("../flights.txt")
	ai.flights = flights
	persons, _ := parsingtxt.Parsing_txt_persons("../persons.txt")
	ai.persons = persons
	ai.current_person_id = len(ai.persons) + 1
	ai.current_flight_id = len(ai.flights) + 1

	return ai
}

func (ai *AdminInterface) Admin() {
	scanner := bufio.NewScanner(os.Stdin)
	var choise int
	for choise != -1 {
		fmt.Println("1. Show users")
		fmt.Println("2. Show flights")
		fmt.Println("3. Add user")
		fmt.Println("4. Add flight")
		fmt.Println("5. Delete user")
		fmt.Println("6. Delete flight")
		fmt.Println("7. Exit")
		fmt.Scan(&choise)
		_ = scanner.Scan()

		switch choise {
		case 1:
			fmt.Println("Persons")
			for _, v := range ai.persons {
				fmt.Println(v)
			}
			_ = scanner.Scan()

		case 2:
			fmt.Println("Flights")
			for _, v := range ai.flights {
				fmt.Println(v)
			}
			fmt.Println("Press Enter to continue...")
			_ = scanner.Scan()

		case 3:
			fmt.Println("Add person")
			id := ai.current_person_id
			ai.current_person_id++

			var surname, name, role string
			var age int

			fmt.Println("Enter surname, name, role")
			_ = scanner.Scan()
			surname = scanner.Text()
			_ = scanner.Scan()
			name = scanner.Text()
			_ = scanner.Scan()
			role = scanner.Text()

			_ = scanner.Scan()
			parsed, err := strconv.Atoi(scanner.Text())
			if err != nil {
				age = 0
			} else {
				age = parsed
			}

			file, err := os.OpenFile("../persons.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("error opening file: %w", err)
				break
			}

			toFile := fmt.Sprintf("\n%d;%s;%s;%s;%d", id, surname, name, role, age)
			_, err = file.WriteString(toFile)
			if err != nil {
				fmt.Println("error writing to file: %w", err)
				break
			}

			file.Close()

			persons, err := parsingtxt.Parsing_txt_persons("../persons.txt")
			if err != nil {
				fmt.Println("2. ", err.Error())
			} else {
				ai.persons = persons
			}

		case 4:
			_ = scanner.Scan()
			fmt.Println("Add flight")

			id := ai.current_flight_id
			ai.current_flight_id++
			var pod, dest, when string
			var dur int

			fmt.Println("Enter pod, dest, when")
			_ = scanner.Scan()
			pod = scanner.Text()
			_ = scanner.Scan()
			dest = scanner.Text()
			_ = scanner.Scan()
			when = scanner.Text()

			_ = scanner.Scan()
			parsed, err := strconv.Atoi(scanner.Text())
			if err != nil {
				dur = 0
			} else {
				dur = parsed
			}

			file, err := os.OpenFile("../flights.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("error opening file: %w", err)
				break
			}

			toFile := fmt.Sprintf("\n%d;%s;%s;%s;%d;", id, pod, dest, when, dur)
			_, err = file.WriteString(toFile)
			if err != nil {
				fmt.Println("error writing to file: %w", err)
				break
			}

			file.Close()

			flights, err := parsingtxt.Parsing_txt_flights("D:\\myFiles\\go_projects\\Aviasales_on_Golang\\flights.txt")
			if err != nil {
				fmt.Println("2. ", err.Error())
			} else {
				ai.flights = flights
			}

		case 5:
			fmt.Println("Delete user")

		case 6:
			fmt.Println("Delete flight")

		case 7:
			fmt.Println("Exit")
			choise = -1

		default:
			fmt.Println("Again")

		}
	}
}
