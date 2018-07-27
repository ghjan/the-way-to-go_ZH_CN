package main

import (
	"sync"
	"time"
	"fmt"
)

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432
Once

Once is an object that will perform exactly one action.

Once 的作用是多次调用但只执行一次，Once 只有一个方法，Once.Do()，向 Do 传入一个函数，
这个函数在第一次执行 Once.Do() 的时候会被调用，以后再执行 Once.Do() 将没有任何动作，即使传入了其它的函数，也不会被执行，
如果要执行其它函数，需要重新创建一个 Once 对象。
 */

func main() {

	o := &sync.Once{}
	go myFun(o)
	go myFun(o)
	time.Sleep(time.Second * 2)
}

func myFun(o *sync.Once) {

	fmt.Println("Begin function")

	o.Do(func() {
		fmt.Println("Working...")
	})

	fmt.Println("Function end")
}
