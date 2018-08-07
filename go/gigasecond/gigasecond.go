// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond returns the moment when someone has lived for a gigasecond (10^9 seconds).
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * time.Duration(math.Pow(10, 9)))
	return t
}
