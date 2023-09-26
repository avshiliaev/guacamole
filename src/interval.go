package src

import "fmt"

// Interval represents a range of integers from start to end.
type Interval struct {
	start int // The start of the interval (inclusive).
	end   int // The end of the interval (inclusive).
}

// NewInterval returns a new Interval with the given start and end values.
func NewInterval(s, e int) *Interval {
	return &Interval{
		start: s,
		end:   e,
	}
}

// String returns a string representation of the Interval in the format "[start, end]".
func (i Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.start, i.end)
}
