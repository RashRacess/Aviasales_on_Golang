package parsingtxt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	f "github.com/RashRacess/Flight"
	p "github.com/RashRacess/Person"
)

func Parsing_txt_persons(filepath string) ([]p.Person, error) {
	persons := make([]p.Person, 0, 10)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.TrimSpace(line), ";")
		if len(parts) != 5 {
			fmt.Printf("Invalid person data format: %s\n", line)
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Invalid person ID: %s\n", parts[0])
			continue
		}

		age, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Printf("Invalid person age: %s\n", parts[4])
			continue
		}

		persons = append(persons, p.CreatePerson(id, parts[1], parts[2], parts[3], age))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return persons, nil
}

func Parsing_txt_flights(filepath string) ([]f.Flight, error) {
	flights := make([]f.Flight, 0, 10)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) != 6 {
			fmt.Printf("Invalid flight data format: %s\n", line)
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Invalid flight ID: %s\n", parts[0])
			continue
		}

		duration, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Printf("Invalid flight duration: %s\n", parts[4])
			continue
		}

		flights = append(flights, f.CreateFlight(id, parts[1], parts[2], parts[3], duration))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return flights, nil
}
