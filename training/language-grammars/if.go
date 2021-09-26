package main

import "fmt"

func  main ()  {
	number := 10
	if 100 > number {
		number += 3
	} else if 100 < number {
		number -= 2
	} else {
		fmt.Println("OK!")
	}

	if number := 4; 100 > number {
		number += 3
	} else if 100 < number {
		number -= 2
	} else {
		fmt.Println("OK!")
	}
}

