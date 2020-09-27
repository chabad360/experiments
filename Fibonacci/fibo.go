package main

// This is an attempt at making different fibonacci implementations.

import (
	"fmt"
)

// The traditional implementation
func fib1(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}

// A slightly more novel approach:
// We put the numbers into processing groups, and deal with them one group at a time.
func fib2(n uint) uint {
	var nums = [][]uint{{n}} // The array that contains arrays of numbers to fibo, pre-initialized with the first number n.
	var final uint           // The final fibo number.

	for i := 0; i < len(nums); i++ { // For each number group to process
		var list []uint
		for ii := 0; ii < len(nums[i]); ii++ {
			switch nums[i][ii] {
			case 1, 0:
				final = final + nums[i][ii] // If the number is a 1 or a 0, add it to the final number.
			default:
				list = append(list, nums[i][ii]-uint(1), nums[i][ii]-uint(2)) // If not, add it to the next group.
			}
		}
		if list != nil {
			nums = append(nums, list)
		}
	}

	return final
}

// Building on the fib2 approach, I've done away with the groups, because they are expensive and unnecessary.
// Now it's much faster and efficient. 
func fib3(n uint) uint {
	var nums = []uint{n} // The array that the numbers to fibo, pre-initialized with the first number n.
	var final uint       // The final fibo number.
	var num uint         // The current number to process.

	for len(nums) != 0 { // Until we run out of numbers to process:
		num, nums = nums[0], nums[1:] // Pop the top of the list.
		switch num {
		case 0:
			break // If it's a zero, skip it.
		case 1:
			final++ // If the number is a 1, add it to the final number.
		default:
			nums = append(nums, num-uint(1), num-uint(2)) // If it's anything else, add it to the proccessing group.
		}
	}

	return final
}

func main() {
	fmt.Println("start")
	n := uint(10)
	fmt.Println(fib3(n))
}
