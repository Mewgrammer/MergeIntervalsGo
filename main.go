package main

import (
	"fmt"
	"github.com/mewgrammer/intervals/intervals"
	"os"
)

func main()  {
	var intervalSlice []intervals.Interval
	if len(os.Args[1]) > 0 {
		result, err := intervals.ParseSlice(os.Args[1])
		if err != nil {
			fmt.Printf("Failed to parse arguments: %v", err)
		}
		intervalSlice = result
	} else {
		intervalSlice = []intervals.Interval {
			{25, 30},
			{2, 19},
			{14,23},
			{4,8},
		}
	}
	result := intervals.Merge(intervalSlice)
	fmt.Printf("Result: %v", result)
}
