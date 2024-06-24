package essence

import "time"

type Flight struct {
	ID    int
	POD   string
	Dect  string
	Time  time.Time
	Costs []int
}

func (f *Flight) GetID() int         { return f.ID }
func (f *Flight) GetPOD() string     { return f.POD }
func (f *Flight) GetDect() string    { return f.Dect }
func (f *Flight) GetTime() time.Time { return f.Time }
func (f *Flight) GetCosts() []int    { return f.Costs }

func (f *Flight) SetID(id int)         { f.ID = id }
func (f *Flight) SetPOD(pod string)    { f.POD = pod }
func (f *Flight) SetDect(dect string)  { f.Dect = dect }
func (f *Flight) SetTime(t time.Time)  { f.Time = t }
func (f *Flight) SetCosts(costs []int) { f.Costs = costs }

func NewFlight(id int, pod, dect string, t time.Time, costs []int) Flight {
	return Flight{
		ID:    id,
		POD:   pod,
		Dect:  dect,
		Time:  t,
		Costs: costs,
	}
}

func NewFlightWithDefaults() Flight {
	return Flight{
		ID:    -1,
		POD:   "undef",
		Dect:  "undef",
		Time:  time.Now(),
		Costs: []int{0, 0, 0, 0},
	}
}
