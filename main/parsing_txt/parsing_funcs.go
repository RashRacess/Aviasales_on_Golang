package parsingtxt

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	f "github.com/RashRacess/Flight"
	p "github.com/RashRacess/Person"
)

func Paring_txt_persons(filepath string) ([]p.Person, error) {
	persons := make([]p.Person, 0, 10)
	file, err := os.Open(filepath)
	if err != nil {
		return []p.Person{}, errors.New("error opening file")
	}
	defer file.Close()

	data := make([]byte, 64)
	var person []string
	for {
		line, err := file.Read(data)
		if err == io.EOF {
			break
		}
		person = strings.Split(string(line), ";")
		if len(person) != 4 {
			continue
		}

		id, err := strconv.Atoi(person[0])
		if err != nil {
			continue
		}
		age, err := strconv.Atoi(person[3])
		if err != nil {
			continue
		}
		persons = append(persons, p.CreatePerson(id, person[1], person[2], age))
	}
	return persons, nil
}

func Paring_txt_flights(filepath string) ([]f.Flight, error) {
	flights := make([]f.Flight, 0, 10)
	file, err := os.Open(filepath)
	if err != nil {
		return []f.Flight{}, errors.New("error opening file")
	}
	defer file.Close()

	data := make([]byte, 64)
	var flight []string
	for {
		line, err := file.Read(data)
		if err == io.EOF {
			break
		}
		flight = strings.Split(string(line), ";")
		if len(flight) != 4 {
			continue
		}

		id, err := strconv.Atoi(flight[0])
		if err != nil {
			continue
		}

		dur, err := strconv.Atoi(flight[5])
		if err != nil {
			continue
		}
		
		flights = append(flights, f.CreateFlight(id, flight[1], flight[2], flight[3], dur))
	}
	return flights, nil
}
