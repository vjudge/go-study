package main

import "fmt"

func main ()  {
	fmt.Println("----------------------")
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		if num % 2 == 1 {
			nums[i] = num * num
		}
	}
	fmt.Println("nums:", nums)
	// nums: [1 2 9 4 25]

	fmt.Println("----------------------")
	// range表达式只会在 for 语句开始执行时被求值一次，无论后边会有多少次迭代
	// range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。
	arys := [...]int{1, 2, 3, 4, 5}
	maxa := len(arys) - 1
	for i, ary := range arys {
		if i == maxa {
			arys[0] = 2 * ary
		} else {
			arys[i + 1] = 2 * ary
		}
	}
	fmt.Println("arys:", arys)
	// arys: [10 2 4 6 8]

	fmt.Println("----------------------")
	// range表达式只会在 for 语句开始执行时被求值一次，无论后边会有多少次迭代
	// range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。
	slice1 := []int{1, 2, 3, 4, 5}
	maxs := len(slice1) - 1
	for i, ele := range slice1 {
		if i == maxs {
			slice1[0] = 2 * ele
		} else {
			slice1[i + 1] = 2 * ele
		}
	}
	fmt.Println("slice1:", slice1)
	// slice1:   [32 2 4 8 16]
}