package essence

import "time"

type Ticket struct {
	ID   int
	POD  string
	Dect string
	Time time.Time
	cost int
}

func (t *Ticket) GetID() int         { return t.ID }
func (t *Ticket) GetPOD() string     { return t.POD }
func (t *Ticket) GetDect() string    { return t.Dect }
func (t *Ticket) GetTime() time.Time { return t.Time }
func (t *Ticket) GetCost() int       { return t.cost }

func (t *Ticket) SetID(id int)           { t.ID = id }
func (t *Ticket) SetPOD(pod string)      { t.POD = pod }
func (t *Ticket) SetDect(dect string)    { t.Dect = dect }
func (t *Ticket) SetTime(time time.Time) { t.Time = time }
func (t *Ticket) SetCost(cost int)       { t.cost = cost }

func NewTicket(id int, pod, dect string, time time.Time, cost int) Ticket {
	return Ticket{
		ID:   id,
		POD:  pod,
		Dect: dect,
		Time: time,
		cost: cost,
	}
}

func NewTicketWithDefaultTime() Ticket {
	return Ticket{
		ID:   -1,
		POD:  "undef",
		Dect: "undef",
		Time: time.Now(),
		cost: -1,
	}
}
