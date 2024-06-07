package ticket

import "fmt"

type Ticket struct {
	id_flight_ int
	pod_       string
	dest_      string
	surname_   string
	name_      string
	when_      string
}

func (t Ticket) GetID() int         { return t.id_flight_ }
func (t Ticket) GetPOD() string     { return t.pod_ }
func (t Ticket) GetDest() string    { return t.dest_ }
func (t Ticket) GetSurname() string { return t.surname_ }
func (t Ticket) GetName() string    { return t.name_ }
func (t Ticket) GetTime() string    { return t.when_ }

func (t *Ticket) SetID(id int)              { t.id_flight_ = id }
func (t *Ticket) SetPOD(pod string)         { t.pod_ = pod }
func (t *Ticket) SetDest(dest string)       { t.dest_ = dest }
func (t *Ticket) SetSurname(surname string) { t.surname_ = surname }
func (t *Ticket) SetName(name string)       { t.name_ = name }
func (t *Ticket) SetTime(when string)       { t.when_ = when }

func CreateTicket(id int, pod string, dest string, surname string, name string, when string) Ticket {
	t := Ticket{}
	t.id_flight_ = id
	t.pod_ = pod
	t.dest_ = dest
	t.surname_ = surname
	t.name_ = name
	t.when_ = when
	return t
}

func (t Ticket) String() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s", t.GetID(), t.GetPOD(), t.GetDest(), t.GetSurname(), t.GetName(), t.GetTime())
}

func (t Ticket) Separated(sep string) string {
	return fmt.Sprint(t.GetID()) + sep + t.GetPOD() + sep + t.GetDest() + sep + t.GetSurname() + sep + t.GetName() + sep + t.GetTime()
}
