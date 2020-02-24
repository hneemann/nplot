package plot

import (
	"github.com/hneemann/plot/vg"
	"time"
)

var _ Ticker = &DenseTimeTicks{}

// DenseTimeTicks creates tick marks as dense as possible
type DenseTimeTicks struct {
	// Format is used to format the date
	Format string

	// Time takes a float64 value and converts it into a time.Time.
	// If nil, UTC unix time is used.
	Time func(t float64) time.Time

	// Float takes a time.Time value and converts it into a float64
	// Must be the inverse of Time
	Float func(t time.Time) float64

	// Axis is required to transform the values from
	// the data coordinate system to the graphic coordinate system
	Axis *Axis
}

type dateModifier func(time time.Time) time.Time

type incrementer struct {
	incr, norm dateModifier
}

var incrementerList = []incrementer{
	{daily(1), normTime},
	{daily(2), normDay},
	{weekly, normDay},
	{twoWeekly, normDay},
	{monthly(1), normDay},
	{monthly(2), normMonth},
	{monthly(3), normMonth},
	{monthly(4), normMonth},
	{monthly(6), normMonth},
	{yearly(1), normMonth},
	{yearly(2), normYear(2)},
	{yearly(5), normYear(5)},
	{yearly(10), normYear(10)},
	{yearly(20), normYear(20)},
}

func normYear(i int) dateModifier {
	return func(t time.Time) time.Time {
		y := (t.Year() / i) * i
		return time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
	}
}

func normMonth(t time.Time) time.Time {
	y := t.Year()
	return time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
}

func normDay(t time.Time) time.Time {
	y := t.Year()
	m := t.Month()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func normTime(t time.Time) time.Time {
	y := t.Year()
	m := t.Month()
	d := t.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func daily(days int) dateModifier {
	return func(t time.Time) time.Time {
		y := t.Year()
		m := t.Month()
		d := t.Day()
		return time.Date(y, m, d+days, 0, 0, 0, 0, t.Location())
	}
}
func weekly(t time.Time) time.Time {
	y := t.Year()
	m := t.Month()
	d := t.Day() + 7
	if d > 28 {
		d = 1
		m += 1
	}
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func twoWeekly(t time.Time) time.Time {
	y := t.Year()
	m := t.Month()
	d := t.Day() + 14
	if d > 28 {
		d = 1
		m += 1
	}
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func monthly(month time.Month) dateModifier {
	return func(t time.Time) time.Time {
		y := t.Year()
		m := t.Month() + month
		d := 1
		return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	}
}

func yearly(years int) func(t time.Time) time.Time {
	return func(t time.Time) time.Time {
		y := t.Year() + years
		return time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
	}
}

func (t *DenseTimeTicks) Ticks(min, max float64, stringSizer StringSizer, axisSize vg.Length) []Tick {
	if t.Time == nil || t.Float == nil {
		t.Time = func(t float64) time.Time {
			return time.Unix(int64(t), 0).In(time.UTC)
		}
		t.Float = func(t time.Time) float64 {
			return float64(t.Unix())
		}
	}

	minTime := t.Time(min)
	size := stringSizer(minTime.Format(t.Format))

	index := 0
	for {
		t0 := incrementerList[index].norm(minTime)
		t1 := incrementerList[index].incr(t0)

		space := vg.Length(t.Axis.Norm(t.Float(t1))-t.Axis.Norm(t.Float(t0))) * axisSize

		if space > size || index == len(incrementerList)-1 {
			break
		}
		index++
	}

	incrementer := incrementerList[index]
	tickTime := incrementer.norm(minTime)

	for t.Float(tickTime) < min {
		tickTime = incrementer.incr(tickTime)
	}

	var ticker []Tick
	for {
		v := t.Float(tickTime)
		if v > max {
			break
		}
		ticker = append(ticker, Tick{
			Value: v,
			Label: t.Time(v).Format(t.Format),
		})
		tickTime = incrementer.incr(tickTime)
	}

	return ticker
}
