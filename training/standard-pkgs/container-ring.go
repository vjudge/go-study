package main

import (
	"container/ring"
	"fmt"
)

func main ()  {
	n := 5
	// 初始化的时候，需要定义好环的大小
	// 创建一个环, 包含 5 个元素
	mring := ring.New(n)
	// 赋值
	for i := 1; i <= n; i ++ {
		mring.Value = i
		mring = mring.Next()
	}
	printRing(mring, n) // 1 2 3 4 5

	fmt.Println("curHead: ", mring.Value)
	mring.Unlink(2)
	printRing(mring, mring.Len())

	// Do 会依次将每个节点的 Value 当作参数调用这个函数 f
	sum := 0
	mring.Do(func(val interface{}) {
		sum += val.(int)
	})
	fmt.Printf("sum ring: %d\n", sum)

}


func printRing (mring *ring.Ring, n int)  {
	for i := 1; i <= n; i++ {
		fmt.Println(mring.Value)
		mring = mring.Next()
	}
}