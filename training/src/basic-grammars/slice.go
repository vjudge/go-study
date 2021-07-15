package main

import "fmt"

func main () {
  // 声明切片
  var slice1 []int
  fmt.Println("slice1 = ", slice1)
  var slice2 = []int{1, 2, 3}
  fmt.Println("slice2 = ", slice2)
  // 短变量
  slice3 := []int{1, 2, 3}
  fmt.Println("slice3 = ", slice3)
   // make()
   var slice4 []int = make([]int, 0)
   fmt.Println("slice4 = ", slice4)
   // 数组
   array := [5]int{1, 2, 3, 4, 5}
   // 前包后不包
   slice5 := array[1:4]
   fmt.Println("slice5 = ", slice5)
   // 切片初始化
   array2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
   slice6 := array2[4:8]
   fmt.Println("slice6 = ", `[4:8]`, slice6)
   slice7 := array2[:5]
   fmt.Println("slice7 = ", `[:5]`, slice7)
   slice8 := array2[6:]
   fmt.Println("slice8 = ", `[6:]`, slice8)
   slice9 := array2[:]
   fmt.Println("slice9 = ", `[:]`, slice9)
   slice10 := array2[:len(array2)-1] //去掉切片的最后一个元素
   fmt.Println("slice10 = ", `:len(array2)-1`, slice10)
}
