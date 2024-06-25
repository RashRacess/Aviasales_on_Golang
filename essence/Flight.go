package essence

import (
	"fmt"
	"strings"
	"time"
)

type Flight struct {
	id_    int
	pod_   string
	dest_  string
	time_  time.Time
	costs_ []int
}

func (f *Flight) GetID() int         { return f.id_ }
func (f *Flight) GetPOD() string     { return f.pod_ }
func (f *Flight) GetDest() string    { return f.dest_ }
func (f *Flight) GetTime() time.Time { return f.time_ }
func (f *Flight) GetCosts() []int    { return f.costs_ }

func (f *Flight) SetID(id int)         { f.id_ = id }
func (f *Flight) SetPOD(pod string)    { f.pod_ = pod }
func (f *Flight) SetDest(dect string)  { f.dest_ = dect }
func (f *Flight) SetTime(t time.Time)  { f.time_ = t }
func (f *Flight) SetCosts(costs []int) { f.costs_ = costs }

func (f *Flight) String() string {
	return fmt.Sprintf("%d - %s - %s - %v - %s", f.GetID(), f.GetPOD(), f.GetDest(), f.GetTime().Format("01-02-2006"), strings.Trim(fmt.Sprint(f.GetCosts()), "[]"))
}

func (f *Flight) StringWithSeparator(separator string) string {
	return fmt.Sprintf("%d%s%s%s%s%s%v%s%s", f.GetID(), separator, f.GetPOD(), separator, f.GetDest(), separator, f.GetTime().Format("01-02-2006"), separator, strings.Trim(fmt.Sprint(f.GetCosts()), "[]"))
}

func NewFlight(id int, pod, dect string, t time.Time, costs []int) Flight {
	return Flight{
		id_:    id,
		pod_:   pod,
		dest_:  dect,
		time_:  t,
		costs_: costs,
	}
}

func NewFlightWithDefaults() Flight {
	return Flight{
		id_:    -1,
		pod_:   "undef",
		dest_:  "undef",
		time_:  time.Now(),
		costs_: []int{0, 0, 0, 0},
	}
}
