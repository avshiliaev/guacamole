package guacamole

import (
	"guacamole/src"
	"reflect"
	"testing"
)

func TestMergeIntervalsFour(t *testing.T) {
	tests := []struct {
		name   string
		input  src.IntervalList
		output src.IntervalList
	}{
		{
			name: "Partially Overlapping",
			input: src.IntervalList{
				src.NewInterval(25, 30),
				src.NewInterval(2, 19),
				src.NewInterval(14, 23),
				src.NewInterval(4, 8),
			},
			output: src.IntervalList{
				src.NewInterval(2, 23),
				src.NewInterval(25, 30),
			},
		},
		{
			name: "Non-Overlapping",
			input: src.IntervalList{
				src.NewInterval(1, 3),
				src.NewInterval(4, 6),
				src.NewInterval(8, 10),
				src.NewInterval(15, 18),
			},
			output: src.IntervalList{
				src.NewInterval(1, 3),
				src.NewInterval(4, 6),
				src.NewInterval(8, 10),
				src.NewInterval(15, 18),
			},
		},
		{
			name: "Completely Overlapping",
			input: src.IntervalList{
				src.NewInterval(1, 10),
				src.NewInterval(4, 6),
				src.NewInterval(3, 7),
				src.NewInterval(8, 9),
			},
			output: src.IntervalList{
				src.NewInterval(1, 10),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.MergeIntervals()
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("got %v, want %v", got, tt.output)
			}
		})
	}
}
