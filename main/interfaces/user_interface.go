package interfaces

import (
	t "aviasales/main/Ticket"
	parsingtxt "aviasales/main/parsing_txt"
	"bufio"
	"fmt"
	"os"
	"strconv"

	f "github.com/RashRacess/Flight"
	p "github.com/RashRacess/Person"
)

type UserInterface struct {
	person  p.Person
	flights []f.Flight
	tickets []t.Ticket
}

func CreateUserInterface(p p.Person) UserInterface {
	ui := UserInterface{}
	ui.person = p
	flights, _ := parsingtxt.Parsing_txt_flights("./parsing_txt/flights.txt")
	ui.flights = flights
	ui.tickets = make([]t.Ticket, 0, 5)
	return ui
}

func (ui UserInterface) User() {
	scanner := bufio.NewScanner(os.Stdin)

	var choise int
	for choise != -1 {
		fmt.Println("1. Show flights")
		fmt.Println("2. Order ticket")
		fmt.Println("3. Show bucket")
		fmt.Println("4. Undo ticket")
		fmt.Println("5. Buy tickets")
		fmt.Println("6. Exit")
		fmt.Scan(&choise)
		_ = scanner.Scan()

		switch choise {
		case 1:
			fmt.Println("Show flights")
			for _, v := range ui.flights {
				fmt.Println(v)
			}
			_ = scanner.Scan()

		case 2:
			var ticket_num int
			var ticket t.Ticket
			fmt.Println("Order ticket")
			fmt.Println("Enter number of ticket you need to order")
			_ = scanner.Scan()
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				ticket_num = 0
			} else {
				ticket_num = num
			}

			for _, v := range ui.flights {
				if v.GetID() == ticket_num {
					ticket = t.CreateTicket(v.GetID(), v.GetPOD(), v.GetDest(), ui.person.GetSurname(), ui.person.GetName(), v.GetTime())
					ui.tickets = append(ui.tickets, ticket)
				}
			}

		case 3:
			fmt.Println("Bucket")
			for _, v := range ui.tickets {
				fmt.Println(v)
			}
		case 4:
		case 5:
		case 6:
			fmt.Println("Exit")
			choise = -1
		}
	}
}
