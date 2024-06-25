package essence

import (
	"fmt"
	"time"
)

type Ticket struct {
	id_   int
	pod_  string
	dest_ string
	time_ time.Time
	cost_ int
}

func (t *Ticket) GetID() int         { return t.id_ }
func (t *Ticket) GetPOD() string     { return t.pod_ }
func (t *Ticket) GetDest() string    { return t.dest_ }
func (t *Ticket) GetTime() time.Time { return t.time_ }
func (t *Ticket) GetCost() int       { return t.cost_ }

func (t *Ticket) SetID(id int)           { t.id_ = id }
func (t *Ticket) SetPOD(pod string)      { t.pod_ = pod }
func (t *Ticket) SetDest(dect string)    { t.dest_ = dect }
func (t *Ticket) SetTime(time time.Time) { t.time_ = time }
func (t *Ticket) SetCost(cost int)       { t.cost_ = cost }

func (t *Ticket) String() string {
	return fmt.Sprintf("%d - %s - %s - %v - %d", t.GetID(), t.GetPOD(), t.GetDest(), t.GetTime().Format("01-02-2006"), t.GetCost())
}

func NewTicket(id int, pod, dect string, time time.Time, cost int) Ticket {
	return Ticket{
		id_:   id,
		pod_:  pod,
		dest_: dect,
		time_: time,
		cost_: cost,
	}
}

func NewTicketWithDefault() Ticket {
	return Ticket{
		id_:   -1,
		pod_:  "undef",
		dest_: "undef",
		time_: time.Now(),
		cost_: -1,
	}
}
