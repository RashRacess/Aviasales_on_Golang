package essence

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

func CreatePerson(id int, surname, name, role string) Person {
	return Person{id, surname, name, role}
}
