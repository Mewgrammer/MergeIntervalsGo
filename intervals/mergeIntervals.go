package intervals

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start, end int
}

func ParseSlice(serialized string) ([]Interval, error ) {
	var intervals []Interval
	var err error = nil
	var errorStr = ""
	if len(serialized) > 0 {
		intervalArgs := strings.Split(serialized, "[")
		for _, intervalStr := range intervalArgs {
			if len(intervalStr) <= 1 {
				continue
			}
			values := strings.Replace(intervalStr, "]", "", 1)
			pair := strings.Split(strings.TrimSpace(values), ",")
			startVal, err := strconv.Atoi(strings.TrimSpace(pair[0]))
			if err != nil {
				errorStr += fmt.Sprintf("%v", err)
			}
			endVal, err := strconv.Atoi(strings.TrimSpace(pair[1]))
			if err != nil {
				errorStr += fmt.Sprintf("%v", err)
			} else {
				interval := Interval{startVal, endVal}
				intervals = append(intervals, interval)
			}
		}
	}
	if len(errorStr) > 0 {
		err = errors.New(errorStr)
	}
	return intervals, err
}

func Merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	// Sort intervals by start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	mergedIntervals := []Interval { intervals[0] }

	for _, val := range intervals[1:]   {
		top := &mergedIntervals[0]
		if  top.end >= val.start {
			top.start = int(math.Min(float64(top.start), float64(val.start)))
			top.end = int(math.Max(float64(top.end), float64(val.end)))
		} else {
			mergedIntervals = append([]Interval{val}, mergedIntervals...) // push to front
		}
	}
	return mergedIntervals
}