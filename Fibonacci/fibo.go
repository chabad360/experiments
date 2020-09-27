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
	var nums [][]uint
	var final uint
	 
	nums = append(nums, []uint{n})

	for i := 0; i < len(nums); i++ {
		var list []uint
		for ii := 0; ii < len(nums[i]); ii++ {
			switch nums[i][ii] {
			case 1:
				final++
			case 0:
				continue
			default:
				list = append(list, nums[i][ii] - 1, nums[i][ii] - 2)
			}
		}
		nums = append(nums, list)
	}

	return final
}

func main() {
	n := uint(10)
	fmt.Println(fib2(n))
}