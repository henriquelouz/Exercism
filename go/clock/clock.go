package clock

import "fmt"

// Clock ...
type Clock struct {
	min int
}

// New ...
func New(hours int, minutes int) Clock {
	c := Clock{(60*hours + minutes) % 1440}
	for c.min < 0 {
		c.min += 1440
	}
	return c
}

// Add ...
func (c Clock) Add(min int) Clock {
	c.min = (c.min + min) % 1440
	for c.min < 0 {
		c.min += 1440
	}
	return c
}

// Subtract ...
func (c Clock) Subtract(min int) Clock {
	c.min = (c.min - min) % 1440
	for c.min < 0 {
		c.min += 1440
	}
	return c
}

// String ...
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}
