package main

import (
    "fmt"
)

// func twoSum(nums []int, target int) (result []int) {
//     for i, outer := range nums {
//         for j, inner := range nums {
//             if i != j && outer + inner == target {
//                 result = []int{i, j}
//                 return result
//             }
//         }
//     }

//     return
// }

func twoSum(nums []int, target int) (result []int) {
    hash := map[int]int{} // key: number, val: index

    for idx, ele := range nums {
        desired := target - ele

        if hashedIdx, ok := hash[desired]; ok {
            result = []int{hashedIdx, idx}
            return
        } else {
           hash[ele] = idx
        }
    }

    return
}

func main() {
    fmt.Println(twoSum([]int{2,7,11,15}, 9))
    fmt.Println(twoSum([]int{3,2,4}, 6))
}
