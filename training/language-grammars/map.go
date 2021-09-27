package main

import "fmt"

func main ()  {
	mmap := map[string]int{
		"first": 1,
		"second": 2,
	}
	fmt.Println(mmap)
	key := "first"
	//key := "three"
	val, ok := mmap[key]
	if ok {
		fmt.Printf("The element of key %s: %d", key, val)
	} else {
		fmt.Println("not found", ok)
	}
}
