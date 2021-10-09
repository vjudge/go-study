package main

import (
	"errors"
	"fmt"
)

func main ()  {
	// 正确捕获 panic
	fmt.Println("start main")
	defer func() {
		fmt.Println("start defer")
		if p := recover(); p != nil {
			fmt.Println("panic: ", p)
		}
		fmt.Println("end defer")
	}()
	panic(errors.New("something is wrong"))
	fmt.Println("end main")
	//// 依然会异常退出
	//fmt.Println("start main")
	//panic(errors.New("something is wrong"))
	//p := recover()
	//fmt.Println("panic: ", p)
	//fmt.Println("end main")
}
