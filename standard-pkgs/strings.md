# strings


### strings.Builder
* 已存在的内容不可变，但可以拼接更多的内容
* 减少了内存分配和内容拷贝的次数
  * 在向Builder值拼接内容的时候并不一定会引起扩容。只要内容容器的容量够用，扩容就不会进行，针对于此的内存分配也不会发生
  * 只要没有扩容，Builder值中已存在的内容就不会再被拷贝
* 可将内容重置，可重用值



### strings.Reader
* 高效读取字符串
* 在读取的过程中，Reader值会保存已读取的字节的计数（以下简称已读计数）



### strings.Compare(s1, s2)
字符串比较，如果相同返回 0，不同返回 -1



### strings.Contains(s, sub)
判断子串。如果是返回 true，如果不是返回 false



### strings.ContainsAny(s, sub)
第二个参数中任意一个字符出现在第一个字符中，则返回true



### strings.Count(s, sub)
计算子串在字符串中出现的无重叠的次数
```go
strings.Count("ydededede", "ded") // 2
```



### strings.Fields(s)
用一个或多个连续的空格分隔字符串 s，返回子字符串的数组(slice)。
如果字符串只包含空格，则返回空列表


### strings.FieldsFunc(s, func(r rune) bool { ... })
用函数指定的条件分隔字符串



### strings.Split(s, sep)
通过 sep 分隔字符串。如果 sep 为空，相当于分成一个个 UTF-8 字符



### strings.SplitAfter(s, sep)
通过 sep 分隔字符串，并且会保留 sep



### strings.SplitN(s, sep, n)
通过 sep分隔字符串，并且返回的结果中最多有 n 个元素
* n > 0 : 返回的 slice 中最多有 n 个元素
* n = 0 : 返回结果为nil
* n < 0 : 正常返回所有的子字符串



### strings.HasPrefix(s, sub)
sub 是否是 s 的前缀。是返回 true，不是返回 false



### strings.HasSuffix(s, sub)
sub 是否是 s 的后缀。是返回 true，不是返回 false



### strings.Index(s, sub)
子串在 s 中第一次出现的位置，如果不存在，返回 -1  
strings.LastIndex(s, sub) 是最后一次出现的位置



### strings.IndexAny(s, sub)
子串中任意一个字符第一次出现在 s 中的位置  
strings.LastIndexAny(s, sub) 是最后一次出现的位置



### strings.IndexFunc(s, func(r rune) bool { ... })
满足函数条件在 s 中第一次出现的位置  
strings.LastIndexFunc(s, sub) 是最后一次出现的位置



### strings.Join(slice, sep)
用分隔符 sep 连接 slice 成字符串



### strings.Repeat(s, n)
将字符串 s 重复 n 次



### strings.Replace(s, oldsub, newsub, n)
将字符串 s 中的所有子串 oldsub，用 newsub 替换 n 次
* n > 0 : 替换 n 次
* n = 0 : 替换 0 次
* n < 0 : 替换所有的子串



### strings.ReplaceAll(s, oldsub, newsub)
将字符串 s 中的所有子串 oldsub，用 newsub 替换



### strings.ToUpper(s)
将字符串转成大写



### strings.ToLower(s)
将字符串转成小写



### strings.Title()
首字母大写，其他字符不处理


### strings.ToTitle()
所有字符大写










