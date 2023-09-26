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

	// Initialize the index and iterate through the sorted list.
	ind := 0
	for i := 1; i < len(l); i++ {
		// If the current interval overlaps with the previous interval, merge them.
		if l[ind].end >= l[i].start {
			l[ind].end = max(l[ind].end, l[i].end)
			l[ind].start = min(l[ind].start, l[i].start)
		} else {
			// Otherwise, add the current interval to the result list
			// and update the index.
			ind++
			l[ind] = l[i]
		}
	}
	// Return the merged IntervalList.
	return l[:ind+1]
}

// compareIntervals compares two intervals by their start values.
func compareIntervals(a, b *Interval) bool {
	return a.start < b.start
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
