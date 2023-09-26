# Overview

Here's a brief overview of the `src` library API.

## Interval

The `Interval` struct represents a range of integers from `start` to `end`. It has the following fields:

- `start` (int): The start of the interval (inclusive).
- `end` (int): The end of the interval (inclusive).

### NewInterval

The `NewInterval` function returns a new `Interval` with the given `start` and `end` values.

```go
func NewInterval(start, end int) *Interval
```

Example usage:

```go
i := NewInterval(1, 5)
fmt.Println(i) // Output: [1, 5]
```

### String

The `String` method returns a string representation of the `Interval` in the format `[start, end]`.

```go
func (i Interval) String() string
```

Example usage:

```go
i := Interval{start: 1, end: 5}
fmt.Println(i.String()) // Output: [1, 5]
```

GitHub Copilot: Sure, here's an example of how to use the `IntervalList` type and `MergeIntervals` function:

## IntervalList

The `IntervalList` type is a slice of `Interval` pointers.

```go
type IntervalList []*Interval
```

### MergeIntervals

The `MergeIntervals` method merges overlapping intervals in the list and returns the result.

```go
func (l IntervalList) MergeIntervals() (result IntervalList)
```

Example usage:

```go
intervals := IntervalList{
    NewInterval(1, 3),
    NewInterval(2, 6),
    NewInterval(8, 10),
    NewInterval(15, 18),
}
merged := intervals.MergeIntervals()
fmt.Println(merged) // Output: [[1, 6], [8, 10], [15, 18]]
```
