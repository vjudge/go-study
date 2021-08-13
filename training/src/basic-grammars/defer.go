package main

import "fmt"

// defer 是先进后出
func main() {
    var whatever [5]struct{}
    for i := range whatever {
        defer fmt.Println(i)
    }
}
// 结果输出
// 4
// 3
// 2
// 1
// 0