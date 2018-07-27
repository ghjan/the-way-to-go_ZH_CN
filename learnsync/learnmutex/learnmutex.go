package main

import (
	"sync"
	"fmt"
)

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432
Locker Mutex

互斥锁用来保证在任一时刻，只能有一个例程访问某对象。Mutex 的初始值为解锁状态。
Mutex 通常作为其它结构体的匿名字段使用，使该结构体具有 Lock 和 Unlock 方法。
 */
type SafeInt struct {
	sync.Mutex
	Num int
}

func main() {
	count := SafeInt{}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			count.Lock()
			count.Num += i
			fmt.Print(count.Num, " ")
			count.Unlock()
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
