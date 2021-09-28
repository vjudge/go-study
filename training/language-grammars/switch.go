package main

import "fmt"

func main ()  {
	nums := [...]int{1, 2, 3, 4, 5, 6}
	sum := 6
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			//switch sum {
			//	case (nums[i] + nums[j]): {
			//		fmt.Printf("[%d, %d]: %d + %d = %d \n", i, j, nums[i], nums[j], sum)
			//	}
			//}
			switch 2 + 4 {
				case nums[i], nums[j]: {
					fmt.Printf("[%d, %d]: %d + %d = %d \n", i, j, nums[i], nums[j], sum)
				}
			}
		}
	}
}
