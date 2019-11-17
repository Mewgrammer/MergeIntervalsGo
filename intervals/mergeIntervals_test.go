package intervals

import (
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tables := []struct {
		input []Interval
		output []Interval
	} {
		{
			[]Interval{{25,30}, {2,19}, {14, 23}, {4, 8} },
			[]Interval{{2,23}, {25, 30},  },
		},
		{
			[]Interval{{1,10}, {5,12} },
			[]Interval{{1,12} },
		},
	}
	for _, table := range tables {
		merged := Merge(table.input)
		for i, mergedInterval := range merged {
			var foundMatch = false
			for _, partialResult := range merged {	// Result may be out of order -> search for match
				if partialResult.start == table.output[i].start && partialResult.end == table.output[i].end {
					foundMatch = true
					break
				}
			}
			if !foundMatch {
				t.Errorf("unexpected result %v - expected: %v", mergedInterval, table.output)
			}
		}
	}
}

func TestParseIntervals(t *testing.T) {
	tables := []struct {
		serialized string
		intervals []Interval
	} {
		{
			"[1,10][5,6]",
			[]Interval{{1,10}, {5,6} } },
		{
			"[ 2,20][5,7]",
			[]Interval{{2,20}, {5,7} } },
		{
			"[3, 30][ 5 , 8 ]",
			[]Interval{{3,30}, {5,8} } },
		{
			"[ 4,40][5,9]",
			[]Interval{{4,40}, {5,9} } },
		{
			"[4,40 ][5,9]",
			[]Interval{{4,40}, {5,9} },
		},
	}
	for _, table := range tables {
		intervals, err := ParseSlice(table.serialized)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i, interval := range intervals {
			if interval.start != table.intervals[i].start || interval.end != table.intervals[i].end {
				t.Errorf("%v is not equal to expected %v", interval, table.intervals[i])
			}
		}
	}
}
