package src

import "sort"

type IntervalList []*Interval

// MergeIntervals merges overlapping intervals in the list and returns the result.

// MergeIntervals merges overlapping intervals in the list and returns the result.
func (l IntervalList) MergeIntervals() (result IntervalList) {

	// Sort the intervals by their start values.
	sort.Slice(l, func(i, j int) bool {
		return compareIntervals(l[i], l[j])
	})

	// Use a stack to merge overlapping intervals.
	stack := make(IntervalList, 0, len(l))

	for _, interval := range l {
		current := interval
		// If the stack is empty or the top interval in the stack does not overlap with the current interval,
		// push the current interval onto the stack.
		if len(stack) == 0 || stack[len(stack)-1].end < current.start {
			stack = append(stack, current)
			// If the top interval in the stack overlaps with the current interval, merge the two intervals by
			// updating the end value of the top interval to the end value of the current interval.
		} else if stack[len(stack)-1].end < current.end {
			stack[len(stack)-1].end = current.end
		}
	}

	// Pop stack into the result list.
	for len(stack) > 0 {
		result = append(result, stack[0])
		stack = stack[1:]
	}
	return
}

// compareIntervals compares two intervals by their start values.
func compareIntervals(a, b *Interval) bool {
	return a.start < b.start
}
