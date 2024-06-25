package interfaces

import (
	"aviasales/essence"
	parsinginformation "aviasales/parsing_information"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	essence.Person
	tickets []essence.Ticket
}


func UserInterface(person essence.Person) {
	fmt.Println("Hello and welcome to the los pollos...")
	
	user := User{person, []essence.Ticket{}}
	var flights = parsinginformation.Parse_txt_flights("../parsing_information/flights.txt")
	scanner := bufio.NewScanner(os.Stdin)
	var isWork bool = true

	for isWork {
		fmt.Println("1. Show profile")
		fmt.Println("2. Show flights")
		fmt.Println("3. Show tickets")
		fmt.Println("4. Buy ticket")
		fmt.Println("5. Undo ticket")
		fmt.Println("6. Exit")

		_ = scanner.Scan()
		choise, err := strconv.Atoi(scanner.Text())
		if err != nil {
			choise = 0
		}
		switch choise {
		case 1:
			fmt.Printf("%s - %s - %s", user.GetName(), user.GetSurname(), user.GetRole())
			_ = scanner.Scan()

		case 2:
			for _, v := range flights {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 3:
			for _, v := range user.tickets {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 4:
			_ = scanner.Scan()
			ticket_num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Wrong flight number")
				_ = scanner.Scan()

				break
			}
			for i, v := range flights {
				if ticket_num == i {
					_ = scanner.Scan()
					class_of_servise, err := strconv.Atoi(scanner.Text())
					if err != nil {
						class_of_servise = 0
					}
					t := essence.NewTicket(0, v.GetPOD(), v.GetDest(), v.GetTime(), class_of_servise)
					user.tickets = append(user.tickets, t)

					break
				}
			}

		case 5:

		case 6:
			fmt.Println("Goodbye")
			isWork = false
		default:
		}
	}
}
