package interfaces

import (
	t "aviasales/main/Ticket"
	parsingtxt "aviasales/main/parsing_txt"
	"bufio"
	"fmt"
	"log"
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

					file, err := os.OpenFile("./parsing_txt/tickets.txt", os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						log.Printf("Ошибка открытия файла: %v", err)
						continue
					}
					defer file.Close()

					_, err = fmt.Fprintf(file, "%s\n", ticket.String())
					if err != nil {
						log.Printf("Ошибка записи в файл: %v", err)
					}
				}
			}

		case 3:
			fmt.Println("Bucket")
			if len(ui.tickets) == 0 {
				fmt.Println("Empty")
			} else {
				for _, v := range ui.tickets {
					fmt.Println(v)
				}
			}
			_ = scanner.Scan()

		case 4:
			fmt.Println("Delete one ticket")
			var id, num int
			fmt.Scan(&id)
			for i, v := range ui.tickets {
				if v.GetID() == id {
					num = i
				}
			}
			for i := num; i < len(ui.tickets)-1; i++ {
				ui.tickets[i] = ui.tickets[i+1]
			}

			ui.tickets = ui.tickets[:len(ui.tickets)-1]

		case 5:
			fmt.Println("You have", len(ui.tickets), "in your bucket")
			fmt.Println("Please, go to nearest aeroport and SAY YOUR NAME for order")
			_ - scanner.Scan()
		case 6:
			fmt.Println("Exit")
			choise = -1
		}
	}
}
