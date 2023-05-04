package main

import (
	"fmt"
	"sort"
)

/*
Problem:
input = array of intervals. each interval has a [start, end]
result = array of non over lapping startend

Merge overlapping intervals!

Constraints:
- 1 <= intervals.length <= 104
- intervals[i].length == 2  (interval array is valid)
- 0 <= starti <= endi <= 104 (interval start and end is valid)

Examples:
Overlap... merge!
|---|
1   3           ->   |------|
  |----|             1      6
  2    6
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
No overlap
|---|
1   3           -> Keep both intervals
       |---|
			 4   7
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Overlap... merge!
|----|
1    4          ->   |----|
 |-|                 1    4
 2 3

Algo:
Sort input intervals by their START

initialize the result array to contain intervals[0]

iterate each currInterval in intervals

    if the prevInterval in result overlaps with currInterval
			create an empty mergedInterval

			mergedInterval START = prevInterval START
			mergedInterval END = larger of prevInterval END and currInterval END

      replace prevInterval with mergedInterval
    else
      append currInterval to result

return result

Analysis:
Sorting takes O(n log(n))
Iterating intervals takes O(n)

O(n log(n) + n) reduces to O(n log(n))
*/

const start, end = 0, 1

func merge(intervals [][]int) [][]int {
	// sort intervals
	sort.Slice(intervals, func(curr, next int) bool {
		return intervals[curr][start] < intervals[next][start]
	})

	result := [][]int{intervals[0]}

	for idx := 1; idx < len(intervals); idx++ {
		prevInterval := result[len(result)-1]
		currInterval := intervals[idx]

		// fmt.Println("prevInterval", prevInterval)
		// fmt.Println("currInterval", currInterval)
		// fmt.Println("")
		if currInterval[start] >= prevInterval[start] && currInterval[start] <= prevInterval[end] {
			newInterval := []int{}
			newInterval = append(newInterval, prevInterval[start])

			// take the largest end for the merged interval
			if currInterval[end] > prevInterval[end] {
				newInterval = append(newInterval, currInterval[end])
			} else {
				newInterval = append(newInterval, prevInterval[end])
			}

			// replace the prev interval with the merged
			result[len(result)-1] = newInterval
		} else {
			result = append(result, currInterval)
		}
	}

	return result
}

func main() {
	// fmt.Println(merge([][]int{{1,4},{1,5}}));
	fmt.Println(merge([][]int{{1, 3}, {8, 10}, {15, 18}, {2, 6}}))
}
