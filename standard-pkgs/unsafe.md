# unsafe



### unsafe.Pointer()
* 可以表示任何指向可寻址的值的指针。
* 可以作为指针值和 uintptr 值质检的桥梁，通过它可以在这两种值之上进行双向的转换。
```go
p := uintptr(unsafe.Pointer(obj))
```


### unsafe.Offsetof()
用于获取两个值在内存中的起始存储地址之间的偏移量，以字节为单位。

















