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
		ClearScreen()
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
			fmt.Println("Your profile")
			fmt.Printf("%s - %s - %s", user.GetName(), user.GetSurname(), user.GetRole())
			_ = scanner.Scan()

		case 2:
			fmt.Println("Available flights")
			for _, v := range flights {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 3:
			fmt.Println("Your tickets")
			for i, v := range user.tickets {
				fmt.Printf("%d - %d - %s - %s - %v - %d", i+1, v.GetID(), v.GetPOD(), v.GetDest(), v.GetTime().Format("01-02-2006"), v.GetCost())
			}
			_ = scanner.Scan()

		case 4:
			fmt.Println("Order ticket")
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
			fmt.Println("undo tickets")
			fmt.Println("Enter ticket's number")

			var number int
			_, err := fmt.Scanf("%d", &number)
			if err != nil {
				fmt.Println("Wrong ticket number:", err)
				return
			}

			if number < 1 || number > len(user.tickets) {
				fmt.Println("Invalid ticket number")
				return
			}

			fmt.Println("Undoing ticket:", user.tickets[number-1])
			user.tickets = append(user.tickets[:number-1], user.tickets[number:]...)

		case 6:
			fmt.Println("Goodbye")
			isWork = false
		default:
		}
	}
}
