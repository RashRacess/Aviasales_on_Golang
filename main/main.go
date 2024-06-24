package main

import (
	parsinginformation "aviasales/parsing_information"
	"fmt"
)

func main() {
	flights := parsinginformation.Parse_txt_flights("../parsing_information/flights.txt")
	for _, v := range flights {
		fmt.Println(v)
	}
}
