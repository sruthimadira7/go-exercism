package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	total := (h*60 + m) % (24 * 60)
	if total < 0 {
		total += 24 * 60
	}
	return Clock{
		hours:   total / 60,
		minutes: total % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}