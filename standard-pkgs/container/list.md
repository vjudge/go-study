# container/list (双向链表)

List 类型由它以及 Element 类型联合表示。

List 根元素永远不会持有任何实际的元素值，该元素值存在就是为了连接这个循环链表的首位两端。

var l list.List 声明的变量，是一个长度为 0 的链表，这个链表持有的根元素将是一共空壳。


### 导入
```go
import "container/list"
```


### list 数据结构
```go
type List struct {
	// 链表跟元素
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	// 链表的长度
	len  int     // current list length excluding (this) sentinel element
}
```


### Element 节点数据结构
```go
// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	// 下一个元素，上一个元素指针
	next, prev *Element

	// The list to which this element belongs.
	// 元素所在的链表
	list *List

	// The value stored with this element.
	// 元素值
	Value interface{}
}
```


### 开箱即用
延迟初始化
* 把初始化操作延后，仅在实际需要的时候才进行。  
* 优点：可以分散初始化操作带来的计算量和存储空间消耗。


### list.New()
初始化一个链表。
```go
mlist = list.New()
```


### list.Front()
获取链表第一个节点。


### list.Back()
获取链表最后一个节点。


### element.Next()
获取链表节点 element 下一个节点。


### element.Prev()
获取链表节点 element 上一个节点。


### list.Len()
返回链表的长度。


### list.PushBack(val)
链表最后插入一个节点。


### list.PushFront(val)
链表头插入一个节点。


### list.Remove(element)
删除某个指定的节点。


### list.InsertBefore(val, element)
在节点 element 前插入 一个节点值。


### list.InsertAfter(val, element)
在节点 element 后插入 一个节点值。


### list.MoveToBack(element)
将节点 element 移动到链表尾部。(必须是链表中的元素)


### list.MoveToFront(element)
将节点 element 移动到链表头部。(必须是链表中的元素)


### list.MoveBefore(element1, element2)
将节点 element1 移动到 element2 之前。(必须是链表中的元素)


### list.MoveAfter(element1, element2)
将节点 element1 移动到 element2 之后。(必须是链表中的元素)


### list.PushBackList(otherlist)
将链表 otherlist 放到链表 list 尾部。


### list.PushFrontList(otherlist)
将链表 otherlist 放到链表 list 头部。






