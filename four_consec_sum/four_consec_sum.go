// Write a function that takes one argument, an array of integers.

// The function should return maximum sum of 4 consecutive numbers in the array.

// If the array contains less than 4 elements, the function should return the lowest negative integer value.

package main

import "fmt"
import "math"

func maximumSum(input []int) (maxSum int) {
	if len(input) < 4 {
		maxSum = math.MinInt32
		return
	}

	maxSum = input[0] + input[1] + input[2] + input[3]
	currSum := maxSum

	for head, tail := 3, 0; head < len(input)-1; {
		head++
		currSum += input[head]
		currSum -= input[tail]
		tail++

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maximumSum([]int{1, 2, 3, 4, 5, 6}))  // 18
	fmt.Println(maximumSum([]int{1, 2, 3}))  // -2147483647
	fmt.Println(maximumSum([]int{3, 2, 6, 5, 1, 2, 9})) //  17
	fmt.Println(maximumSum([]int{-1, -5, -3, 0, -1, 2, -4})) // -2
}
