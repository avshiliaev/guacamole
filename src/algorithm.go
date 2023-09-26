package src

type IntervalList []*Interval

// MergeIntervals merges overlapping intervals in the list and returns the result.
func (l IntervalList) MergeIntervals() (result IntervalList) {
	n := len(l)
	for i := 0; i < n; i++ {
		merged := false
		for j := 0; j < len(result); j++ {
			if ok, interval := merge(l[i], result[j]); ok {
				result[j] = interval
				merged = true
				break
			}
		}
		if !merged {
			result = append(result, l[i])
		}
	}
	return
}

// merge merges two intervals if they overlap and returns true and the merged interval.
// Otherwise, it returns false and nil.
func merge(a, b *Interval) (bool, *Interval) {
	if a.end < b.start {
		return false, nil
	}
	if a.start > b.end {
		return false, nil
	}
	return true, NewInterval(min(a.start, b.start), max(a.end, b.end))
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
