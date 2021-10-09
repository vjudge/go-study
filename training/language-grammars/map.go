package main

import (
	"fmt"
)

func main ()  {
	mmap := map[string]int{
		"first": 1,
		"second": 2,
	}
	fmt.Println(mmap)

	// 初始化 0 值问题
	mmap2 := map[int]int{ 2: 0 }
	fmt.Println("mmap2: 1: ", mmap2[1])
	fmt.Println("mmap2: 2: ", mmap2[2])
	key := 1
	val, ok := mmap2[key]
	if ok {
		fmt.Printf("The element of key %s: %d\n", key, val)
	} else {
		fmt.Printf("key %d not found \n", key)
	}

	mmap3 := map[int]func(num int, key int) int {}
	mmap3[1] = func(num int, key int) int {
		return key * num
	}
	mmap3[2] = func(num int, key int) int {
		return key * num
	}
	mmap3[3] = func(num int, key int) int {
		return key * num
	}
	fmt.Println("mmap3[1](2):", mmap3[1](2, 1))
	fmt.Println("mmap3[2](2):", mmap3[2](2, 2))
	fmt.Println("mmap3[3](2):", mmap3[3](2, 3))

	// 实现 Set：map[type]bool
	myset := map[int]bool{}
	myset[1] = true
	if myset[2] {
		fmt.Printf("this key %d is exist", key)
	} else {
		fmt.Printf("this key %d is not exist", key)
	}
	myset[5] = true
	delete(myset, 1)
	fmt.Println("myset:", myset)
}


