package src

import "sort"

// IntervalList defines a type for a list of intervals
type IntervalList []*Interval

// MergeIntervals merges overlapping intervals in the list and returns the result.
func (l IntervalList) MergeIntervals() (result IntervalList) {

	// Sort the intervals by their start values.
	sort.Slice(l, func(i, j int) bool {
		return compareIntervals(l[i], l[j])
	})

	// Initialize the current interval to the first interval in the sorted list.
	current := l[0]
	// Iterate over the remaining intervals in the sorted list.
	for i := 1; i < len(l); i++ {
		// If the current interval overlaps with the next interval.
		if current.end >= l[i].start {
			// Merge the two intervals by updating the end value of the current interval.
			current.end = max(current.end, l[i].end)
		} else {
			// Otherwise, add the current interval to the result list
			// and update the current interval to the next interval.
			result = append(result, current)
			current = l[i]
		}
	}
	// Add the final current interval to the result list.
	result = append(result, current)
	return
}

// compareIntervals compares two intervals by their start values.
func compareIntervals(a, b *Interval) bool {
	return a.start < b.start
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
