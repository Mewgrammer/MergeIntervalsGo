package main

import (
	"fmt"
	"math"
	"sort"
)
type Interval struct {
	start, end int
}

func main() {
	intervals := []Interval{
		{25, 30},
		{2, 19},
		{14,23},
		{4,8},
	}
	mergedIntervals := mergeIntervals(intervals)
	fmt.Printf("Result: %v", mergedIntervals)
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	// Sort intervals by start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	mergedIntervals := []Interval { intervals[0] }

	for i := 1; i < len(intervals); i++  {
		top := mergedIntervals[0]
		if  top.end >= intervals[i].start {
			top.start = int(math.Min(float64(top.start), float64(intervals[i].start)))
			top.end = int(math.Max(float64(top.end), float64(intervals[i].end)))
		} else {
			mergedIntervals = append([]Interval{intervals[i]}, mergedIntervals...) // push to front
		}
	}
	return mergedIntervals
}