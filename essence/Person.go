package essence

import "fmt"

type Person struct {
	id_      int
	surname_ string
	name_    string
	role_    string
}

func (p Person) GetID() int         { return p.id_ }
func (p Person) GetSurname() string { return p.surname_ }
func (p Person) GetName() string    { return p.name_ }
func (p Person) GetRole() string    { return p.role_ }

func (p *Person) SetId(id int)              { p.id_ = id }
func (p *Person) SetSurname(surname string) { p.surname_ = surname }
func (p *Person) SetName(name string)       { p.name_ = name }
func (p *Person) SetRole(role string)       { p.role_ = role }

func (p *Person) String() string {
	return fmt.Sprintf("%d - %s - %s - %s", p.GetID(), p.GetSurname(), p.GetName(), p.GetRole())
}

func (p *Person) StringWithSeparator(separator string) string {
	return fmt.Sprintf("%d%s%s%s%s%s%s", p.GetID(), separator, p.GetSurname(), separator, p.GetName(), separator, p.GetRole())
}

func NewPerson(id int, surname, name, role string) Person {
	return Person{
		id_:      id,
		surname_: surname,
		name_:    name,
		role_:    role,
	}
}

func NewPersonWithDefaults() Person {
	return Person{
		id_:      -1,
		surname_: "undef",
		name_:    "undef",
		role_:    "undef",
	}
}
