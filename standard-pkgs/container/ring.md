# container/ring
循环链表（环）。

在创建并初始化一个 Ring 时，可以指定它包含的元素数量。 循环链表一旦被创建，其长度是不可变的。

var r ring.Ring 声明的 r 是一共长度为 1 的循环链表。



### 导入
```go
import "container/ring"
```


### 数据结构
```go
// A Ring is an element of a circular list, or ring.
// Rings do not have a beginning or end; a pointer to any ring element
// serves as reference to the entire ring. Empty rings are represented
// as nil Ring pointers. The zero value for a Ring is a one-element
// ring with a nil Value.
//
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}
```


### ring.New(len)
初始化一个循环链表，需要指定环的长度。
```go
mring = ring.New(5)
```


### ring.Next()
获取循环链表下一个节点。


### ring.Prev()
获取循环链表上一个节点。


### ring.Move(n)
指针从当前元素开始向后移动 n 位，负数代表向前移动 n。


### ring.Link(ring1)
将两个 ring 连接到一起 (ring 不能为空)。


### ring.Unlink(n)
从当前元素下一个开始，删除 n 个元素。


### ring.Do(f func(interface{}))
Do 会依次将每个节点的 Value 当作参数调用这个函数 f, 实际上这是策略方法的引用，通过传递不同的函数以在同一个 ring 上实现多种不同的操作。

































