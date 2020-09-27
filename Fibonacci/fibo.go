package main
// This is an attempt at making different fibonacci implementations.

import (
	"fmt"
)

// The traditional implementation
func fib(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fib2(n uint) uint {
	var nums = [][]uint{{n}} // The array that contains arrays of numbers to fibo, pre-initialized with the first number n.
	var final uint // The final fibo number.

	for i := 0; i < len(nums); i++ { // For each number group to process
		var list []uint
		for ii := 0; ii < len(nums[i]); ii++ {
			switch nums[i][ii] {
			case 1, 0:
				final = final + nums[i][ii] // If the number is a 1 or a 0, add it to the final number.
			default:
				list = append(list, nums[i][ii] - uint(1), nums[i][ii] - uint(2)) // If not, add it to the next group.
			}
		}
		if list != nil {
			nums = append(nums, list)
		}
	}

	return final
}

func main() {
	fmt.Println("start")
	n := uint(10)
	fmt.Println(fib2(n))
}