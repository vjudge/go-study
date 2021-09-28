# 结构体和方法

Go语言的结构体类型（Struct）可以封装属性和操作。 


### 声明
结构体类型的字面量由关键字type、类型名称、关键字struct，以及由花括号包裹的若干字段声明组成。其中，每个字段声明独占一行并由字段名称（可选）和字段类型组成。
```
type Person struct {
    Name   string
    Gender string
    Age    uint8
}
```


### 初始化
```
Person{Name: "Robert", Gender: "Male", Age: 33}  
// 如果这里的键值对的顺序与其类型中的字段声明完全相同的话，我们还可以统一省略掉所有字段的名称。
Person{"Robert", "Male", 33} 
```
我们在编写某个结构体类型的值字面量时可以只对它的部分字段赋值，甚至不对它的任何字段赋值。这时，未被显式赋值的字段的值则为其类型的零值。注意，在上述两种情况下，字段的名称是不能被省略的。



### 匿名结构体
在编写匿名结构体的时候需要先写明其类型特征（包含若干字段声明），再写出它的值初始化部分。匿名结构体最大的用处就是在内部临时创建一个结构以封装数据，而不必正式为其声明相关规则。
```
p := struct {
    Name   string
    Gender string
    Age    uint8
}{"Robert", "Male", 33}
```



### 结构体方法
结构体类型可以拥有若干方法（匿名结构体不可能拥有方法）。方法的特殊在于它的声明包含了一个接收者声明。
```
// Person类型的指针方法
func (person *Person) Grow() {
    person.Age++
} 
```



### 委托
Go 语言规范规定，如果一个字段的声明中只有字段的类型名而没有字段的名称，那么它就是一共嵌入字段，也可以称为匿名字段。
```go
type Animal struct { 
	Name string // 学名
	AnimalCategory // 动物基本分类
}

type AnimalCategory struct {
    kingdom string // 界
    phylum string // 门
    class string // 纲
	order string // 目
	family string // 科
    genus string // 属
	species string // 种
}
```































