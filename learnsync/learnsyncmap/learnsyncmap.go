package main

import (
	"sync"
	"fmt"
	"time"
)

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432
Map

并发安全的map，关于golang中map的并发不安全之后再详细介绍。
 */
func main() {
	var m sync.Map
	for i := 0; i < 3; i++ {
		go func(i int) {
			for j := 0; ; j++ {
				m.Store(i, j)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		m.Range(func(key, value interface{}) bool {
			fmt.Printf("%d: %d\t", key, value)
			return true
		})
		fmt.Println()
		time.Sleep(time.Second)
	}
}
