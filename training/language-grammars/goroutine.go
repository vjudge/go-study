package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main ()  {
	var count uint32 = 1
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	// 不会有任何值打印出来
	for i := uint32(1); i <= 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println("i=", i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(11, func() {})
}

//func main ()  {
//	// 不会有任何值打印出来
//	for i := 1; i <= 10; i++ {
//		go func(i int) {
//			fmt.Println("i=", i)
//		}(i)
//	}
//	// 不会有任何值打印出来
//
//	time.Sleep(time.Microsecond * 500)
//	// 各协程执行顺序无法确定
//	// 执行顺序无法确定
//}
