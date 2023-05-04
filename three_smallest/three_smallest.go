package main

import (
	"fmt"
)

/*
smallest = 0

while smallest < last index
	smaller = smallest + 1
	largest = last index

	while smallest < largest
	  complement = 0 - nums[smallest] + nums[smaller]

		if nums[largest] > complement
		  smaller++

		if nums[largest] == complement
		  append to result

		if nums[largest] < complement
		  largest++
*/

func threeSum(nums []int) (result [][]int) {
	for smallest := 0; smallest < len(nums); smallest++ {
		smaller := smallest + 1
		largest := len(nums)-1

		fmt.Println(smallest, smaller, largest)

		for ; smaller < largest; {
			complement := 0 - nums[smallest] + nums[smaller]

			if nums[largest] > complement {
				smaller++
					fmt.Println(smallest, smaller, largest)
				continue
			}

			if nums[largest] == complement {
				result = append(result, []int{nums[smallest], nums[smaller], nums[largest]})
				fmt.Println(smallest, smaller, largest)
				continue
			}

			if nums[largest] < complement {
				largest--
				fmt.Println(smallest, smaller, largest)
				continue
			}
		}
	}

	return
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}
