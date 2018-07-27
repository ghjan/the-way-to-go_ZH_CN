package main

import (
	"sync"
	"fmt"
)

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432
Pool

Pool 用于存储临时对象，它将使用完毕的对象存入对象池中，在需要的时候取出来重复使用，
目的是为了避免重复创建相同的对象造成 GC 负担过重。其中存放的临时对象随时可能被 GC 回收掉（如果该对象不再被其它变量引用）。

从 Pool 中取出对象时，如果 Pool 中没有对象，将返回 nil，但是如果给 Pool.New 字段指定了一个函数的话，
Pool 将使用该函数创建一个新对象返回。

Pool 可以安全的在多个例程中并行使用，但 Pool 并不适用于所有空闲对象，Pool 应该用来管理并发的例程共享的临时对象，
而不应该管理短寿命对象中的临时对象，因为这种情况下内存不能很好的分配，这些短寿命对象应该自己实现空闲列表。

切记，Pool 在开始使用之后，不能再被复制！！！！
 */
func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(100)
	b := p.Get().(int)
	fmt.Println(a, b)
	c := p.Get().(int)
	fmt.Println(c)
	p.Put(101)
	d := p.Get().(int)
	fmt.Println(d)
	p.Put(b)
	p.Put(d)
	e := p.Get().(int)
	fmt.Println(e)
	f := p.Get().(int)
	fmt.Println(f)
}
