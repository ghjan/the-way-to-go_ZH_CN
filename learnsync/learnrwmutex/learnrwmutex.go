package main

import (
	"sync"
	"time"
)

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432
RWMutex

RWMutex 比 Mutex 多了一个“读锁定”和“读写解锁”，可以让多个例程同时读取某对象。

RWMutex 可以安全的在多个例程中并行使用。
 */
var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	go read(1)
	go read(2)

	time.Sleep(2 * time.Second)
}

func read(i int) {
	println(i, "start")

	m.RLock()
	println(i, "reading......")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "end")
}
