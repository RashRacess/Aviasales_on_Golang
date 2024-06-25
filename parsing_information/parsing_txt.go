package parsinginformation

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"aviasales/essence"
)

func Parse_txt_persons(filename string) []essence.Person {
	file, err := os.Open(filename)
	if err != nil {
		return []essence.Person{}
	}
	_ = file
	persons := make([]essence.Person, 0, 15)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		words := strings.Split(ln, ";")

		id, err := strconv.Atoi(words[0])
		if err != nil {
			id = 0
		}
		persons = append(persons, essence.NewPerson(id, words[1], words[2], words[3]))
	}
	return persons
}

func Parse_txt_flights(filename string) []essence.Flight {
	file, err := os.Open(filename)
	if err != nil {
		return []essence.Flight{}
	}
	flights := make([]essence.Flight, 0, 15)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		words := strings.Split(ln, ";")
		id, err := strconv.Atoi(words[0])
		if err != nil {
			id = 0
		}
		t, err := time.Parse("2006-01-02", words[3])
		if err != nil {
			t = time.Now()
		}

		costs_s := strings.Split(words[4], " ")
		costs_i := make([]int, 4)
		for i, v := range costs_s {
			c, err := strconv.Atoi(v)
			if err != nil {
				costs_i[i] = -1
			} else {
				costs_i[i] = c
			}
		}

		flights = append(flights, essence.NewFlight(id, words[1], words[2], t, costs_i))
	}
	return flights
}

func RecordPersonToFile(filename string, data []essence.Person, separator string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	for _, v := range data {
		_, err := file.WriteString(v.StringWithSeparator(separator) + "\n")
		if err != nil {
			continue
		}
	}
	return nil
}

func RecordFlightToFile(filename string, data []essence.Flight, separator string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	for _, v := range data {
		_, err := file.WriteString(v.StringWithSeparator(separator) + "\n")
		if err != nil {
			continue
		}
	}
	return nil
}
