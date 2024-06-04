package authentification

import (
	"errors"

	p "github.com/RashRacess/Person"
)

func Entering(person p.Person, persons *[]p.Person) (p.Person, error) {
	for _, v := range *persons {
		if v.GetSurname() == person.GetSurname() && v.GetName() == person.GetName() {
			return v, nil
		}
	}
	return p.Person{}, errors.New("error entering")
}

func Registration(person p.Person, persons *[]p.Person) {
	if _, err := Entering(person, persons); err != nil {
		return
	} else {
		*persons = append(*persons, person)
	}
}
