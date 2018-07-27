package main

import (
	"time"
	"sync"
	"fmt"
)

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432

WaitGroup

WaitGroup 用于等待一组例程的结束。主例程在创建每个子例程的时候先调用 Add 增加等待计数，
每个子例程在结束时调用 Done 减少例程计数。之后，主例程通过 Wait 方法开始等待，直到计数器归零才继续执行。
 */
func doSomething(milliSecs time.Duration, wg *sync.WaitGroup) {
	duration := milliSecs * time.Millisecond
	time.Sleep(duration)
	fmt.Println("Function in background, duration:", duration)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go doSomething(200, &wg)
	go doSomething(400, &wg)
	go doSomething(150, &wg)
	go doSomething(600, &wg)

	wg.Wait()
	fmt.Println("Done")
}
