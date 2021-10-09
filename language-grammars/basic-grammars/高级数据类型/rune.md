# rune

Go 语言中的标识符可以包含“任何 Unicode 编码可以表示的字母字符”

Go 语言的代码是由 Unicode 字符组成的。Go 语言的所有源代码，都必须按照 Unicode 编码规范中的 UTF-8 编码格式进行编码。



### Unicode
* Unicode 编码规范通常使用十六进制表示法来表示 Unicode 代码点的整数值，并使用“U+”作为前缀
* 在 Unicode 编码规范中，一个字符能且只能由与它对应的那个代码点表示
* Unicode 编码规范提供了三种不同的编码格式，即：UTF-8 (1个字节)、UTF-16 (2个字节) 和 UTF-32 (4个字节)。(Universal Character Set Transformation Format)
  * UTF-8 是一种可变宽的编码方案，与标准的 ASCII 编码是完全兼容的



### string
一个string类型的值是由一系列相对应的 Unicode 代码点的 UTF-8 编码值来表达的

在 Go 语言中，一个string类型的值既可以被拆分为一个包含多个字符的序列，也可以被拆分为一个包含多个字节的序列
* 字符序列：以由一个以rune为元素类型的切片来表示
* 字节序列：可以由一个以byte为元素类型的切片代表


### rune
rune 是 Go 语言特有的一个基本数据类型，它的一个值代表一个字符，即一个 Unicode 字符

当一个string类型的值被转换为[]rune类型值的时候，其中的字符串会被拆分成一个一个的 Unicode 字符
* 每一个字符都会独立成为一个 rune 类型的元素值

#### 声明
```go
// rune 类型实际上就是 int32 类型的一个别名类型
// 一个 rune 类型的值由 4 个字节存储
type rune = init32
```







