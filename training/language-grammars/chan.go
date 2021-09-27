package main

import "fmt"

func main ()  {
	ch1 := make(chan int, 5)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5
	//ch1 <- 6
	for i := 1; i <= 6; i++ {
		ele := <-ch1
		fmt.Println("ele:", i, ele)
	}
}
