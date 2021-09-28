package main

import "fmt"

func main ()  {
	// 定义
	type Person struct {
		Name string
		Gender string
		Age uint8
	}
	// 初始化
	p1 := Person{Name: "Robert", Age: 33}
	fmt.Printf("Name: %s, Gender: %s, Age: %d \n", p1.Name, p1.Gender, p1.Age)
	// Name: Robert, Gender: , Age: 33
	p2 := Person{"vjudge", "Female", 18}
	fmt.Printf("Name: %s, Gender: %s, Age: %d \n", p2.Name, p2.Gender, p2.Age)
	// Name: vjudge, Gender: Female, Age: 18
}