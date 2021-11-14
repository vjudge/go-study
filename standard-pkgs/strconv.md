# strconv
字符串和基本数据类型之间的转换

基本数据类型包括
* 布尔
* 整型（包括有 / 无符号、二进制、八进制、十进制和十六进制）
* 浮点型


## 字符串转整型
### strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
转为有符号整型
* base : 代表字符串按照给定的进制进行解释。base 的取值为 2~36
  * base 值为 0，则会根据字符串的前缀来确定 base 的值
    * "0x" : 表示 16 进制
    * "0" : 表示 8 进制
    * 否则就是 10 进制
* bitSize 表示的是整数取值范围，或者说整数的具体类型
  * 0 : int
  * 8 : int8
  * 16 : int16
  * 32 : int32
  * 64 : int64


### strconv.ParseUint(s string, base int, bitSize int) (n uint64, err error)
转为无符号整型


### strconv.Atoi(s string) (i int64, err error)
是 ParseInt 的便捷版，内部通过调用 ParseInt(s, 10, 0) 来实现的


### strconv.AppendInt(dst []byte, i int64, base int)
将整数转为字符数组， append 到目标字符数组中


### strconv.AppendUint(dst []byte, i uint64, base int)
将无符整数转为字符数组， append 到目标字符数组中



### 整型转为字符串
### strconv.FormatUint(i uint64, base int) string    
无符号整型转字符串  

将整数每一位数字对应到相应的字符，存入字符数组中，最后字符数组转为字符串即为结果


### strconv.FormatInt(i int64, base int) string    
有符号整型转字符串


### strconv.Itoa(i int) string
Itoa 内部直接调用 FormatInt(i, 10) 实现的。base 参数可以取 2~36（0-9，a-z）



## 字符串转布尔值
### strconv.ParseBool(str string) (value bool, err error)
接受字符串: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False


### strconv.FormatBool(b bool) string
直接返回 "true" 或 "false"


### strconv.AppendBool(dst []byte, b bool)
将 "true" 或 "false" append 到 dst 中



## 字符串转浮点数
### strconv.ParseFloat(s string, bitSize int) (f float64, err error)

### strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string
prec 表示有效数字

### strconv.AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int)




## 错误处理
在返回错误的时候，不是直接将下面的变量值返回，是通过构造一个 NumError 类型的 error 对象返回

### ErrRange
表示值超过了类型能表示的最大范围

### ErrSyntax
表示语法错误




