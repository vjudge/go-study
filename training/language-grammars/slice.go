package main

import "fmt"

func main () {
    // 声明切片
    var slice1 []int
    fmt.Println("slice1 = ", slice1) // slice1 =  []

    // 对应底层数组中的第 1 个元素
    var slice2 = []int{1, 2, 3}
    fmt.Println("slice2 = ", slice2) // slice2 =  [1 2 3]
    fmt.Println("slice2 = ", len(slice2), cap(slice2)) // slice2 =  3 3

    // 短变量
    slice3 := []int{1, 2, 3}
    fmt.Println("slice3 = ", slice3) // slice2 =  [1 2 3]

    // make()
    // var slice []type = make([]type, len, cap) // 省略 cap，相当于 cap = len
    var slice4 []int = make([]int, 2, 10)
    fmt.Println("slice4 = ", slice4) // // slice4 = [0, 0]
    fmt.Println("slice4 = ", len(slice4), cap(slice4)) // slice4 =  2 10

    // 数组
    array := [5]int{1, 2, 3, 4, 5}
    // 左闭右开 [1, 4)
    slice5 := array[1:4]
    fmt.Println("slice5 = ", slice5) // slice5 =  [2 3 4]

    // 切片初始化
    array2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    slice6 := array2[4:6]
    fmt.Println("slice6 = ", `[4:6]`, slice6) // slice6 =  [4:6] [4 5]
    slice22 := append(slice6, 111)
    fmt.Println("slice22 = ", slice22) // slice22 =  [4 5 111]
    fmt.Println("slice6 = ", `[4:6]`, slice6) // slice6 =  [4:6] [4 5]
    fmt.Println("slice6 = ", len(slice6), cap(slice6)) // slice2 =  2 7

    slice7 := array2[:5]
    fmt.Println("slice7 = ", `[:5]`, slice7) // slice7 =  [:5] [0 1 2 3 4]

    slice8 := array2[6:]
    fmt.Println("slice8 = ", `[6:]`, slice8) // slice8 =  [6:] [6 7 8 9 10]

    slice9 := array2[:]
    fmt.Println("slice9 = ", `[:]`, slice9) // slice9 =  [:] [0 1 2 3 4 5 6 7 8 9 10]

    //去掉切片的最后一个元素
    slice10 := array2[:len(array2)-1]
    fmt.Println("slice10 = ", `:len(array2)-1`, slice10) // slice10 =  :len(array2)-1 [0 1 2 3 4 5 6 7 8 9]

    // append()
    slice11 := append(slice5, 888)
    fmt.Println("slice11 = ", slice11) // slice11 =  [2 3 4 888]
    // 遍历
    for i, v := range slice6 {
        fmt.Printf("遍历: i = %v , v = %v\n", i, v)
    }
    // 遍历: i = 0 , v = 4
    // 遍历: i = 1 , v = 5
    // 遍历: i = 2 , v = 6
    // 遍历: i = 3 , v = 7

}
