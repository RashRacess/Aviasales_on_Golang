package interfaces

import (
	"aviasales/essence"
	parsinginformation "aviasales/parsing_information"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
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
		ClearScreen()
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
			fmt.Println("Your profile")
			fmt.Printf("%s - %s - %s", admin.GetName(), admin.GetSurname(), admin.GetRole())
			_ = scanner.Scan()

		case 2:
			fmt.Println("Showing users")
			for _, v := range users {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 3:
			fmt.Println("Deleting users")
			fmt.Println("Enter user's number")

			var number int
			_, err := fmt.Scanf("%d", &number)
			if err != nil {
				fmt.Println("Wrong user number:", err)
				return
			}

			index := -1
			for i, v := range users {
				if v.GetID() == number {
					index = i
					break
				}
			}

			if index == -1 {
				fmt.Println("User not found")
				return
			}

			fmt.Println("Deleting user:", users[index])
			users = append(users[:index], users[index+1:]...)
			parsinginformation.RecordPersonToFile("../parsing_information/persons.txt", users, ";")

		case 4:
			notFound := false
			fmt.Println("Adding users")
			fmt.Println("Enter id")
			_ = scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Wrong id")
				break
			}

			for _, v := range users {
				if v.GetID() == id {
					fmt.Println("Wrong id")
					notFound = true
					break
				}
			}
			if notFound {
				break
			}

			fmt.Println("Enter surname")
			_ = scanner.Scan()
			surname := scanner.Text()

			fmt.Println("Enter name")
			_ = scanner.Scan()
			name := scanner.Text()

			fmt.Println("Enter role")
			_ = scanner.Scan()
			role := scanner.Text()

			users = append(users, essence.NewPerson(id, surname, name, role))
			parsinginformation.RecordPersonToFile("../parsing_information/persons.txt", users, ";")

		case 5:
			fmt.Println("Showing flights")
			for _, v := range flights {
				fmt.Println(v.String())
			}
			_ = scanner.Scan()

		case 6:
			fmt.Println("Deleting flights")
			fmt.Println("Enter flight's number")
			_ = scanner.Scan()
			number_s := scanner.Text()
			number, err := strconv.Atoi(number_s)
			if err != nil {
				fmt.Println("Wrong flight number")
				break
			}
			var number_for_deleting int
			for i, v := range flights {
				if v.GetID() == number {
					number_for_deleting = i
				}
			}

			temp := flights[number_for_deleting]
			flights[number_for_deleting] = flights[len(flights)-1]
			flights[len(flights)-1] = temp
			flights = flights[:len(flights)-1]

			parsinginformation.RecordFlightToFile("../parsing_information/flights.txt", flights, ";")

		case 7:
			fmt.Print("Enter flight ID: ")
			_ = scanner.Scan()
			id_s := scanner.Text()
			id, err := strconv.Atoi(id_s)
			if err != nil {
				fmt.Println("Wrong id")
				break
			}

			for _, v := range flights {
				if v.GetID() == id {
					fmt.Println("Wrong id")
					break
				}
			}

			fmt.Print("Enter POD: ")
			_ = scanner.Scan()
			pod := scanner.Text()

			fmt.Print("Enter destination: ")
			_ = scanner.Scan()
			dest := scanner.Text()

			fmt.Print("Enter flight date (MM-DD-YYYY): ")
			_ = scanner.Scan()
			date_s := scanner.Text()
			date, err := time.Parse("2006-01-02", date_s)
			if err != nil {
				date = time.Now()
			}

			fmt.Println("Enter costs(space separated)")
			_ = scanner.Scan()
			costs, _ := parseCosts(scanner.Text(), " ")

			flights = append(flights, essence.NewFlight(id, pod, dest, date, costs))
			parsinginformation.RecordFlightToFile("../parsing_information/flights.txt", flights, ";")

		case 8:
			fmt.Println("Goodbye")
			isWork = false

		default:
			fmt.Println("Try againg")
		}
	}
}

func parseCosts(costsStr string, separator string) ([]int, error) {
	costs := strings.Split(costsStr, separator)
	costsInt := make([]int, len(costs))
	for i, cost := range costs {
		c, err := strconv.Atoi(strings.TrimSpace(cost))
		if err != nil {
			return nil, err
		}
		costsInt[i] = c
	}
	return costsInt, nil
}

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		// Для Linux и macOS
		cmd := exec.Command("/usr/bin/clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		// Для Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		// Если не удается определить операционную систему, ничего не делаем
		return
	}
}
