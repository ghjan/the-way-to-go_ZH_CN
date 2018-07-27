package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Cond

Cond用于在并发环境下routine的等待和通知
 */
func main() {
	m := sync.Mutex{}
	m.Lock() //先锁住
	c := sync.NewCond(&m)
	go func() {
		m.Lock() //请求锁
		defer m.Unlock()
		fmt.Println("3. goroutine is owner of lock")
		time.Sleep(2 * time.Second)
		fmt.Println("4. goroutine will release lock soon (deffered Unlock)")
		/*Broadcase、Signal
			唤醒因wait condition而挂起goroutine，区别是Signal只唤醒一个，而Broadcast唤醒所有。
			允许调用者获取基础锁Locker之后再调用唤醒，但非必需。
		*/
		c.Signal()
		//c.Broadcast() //唤醒所有等待的 Wait
	}()
	fmt.Println("1. main goroutine is owner of lock")
	time.Sleep(1 * time.Second)
	fmt.Println("2. main goroutine is still lockek")
	/*
	必须获取该锁之后才能调用Wait()方法
	Wait方法在调用时会释放底层锁Locker，并且将当前goroutine挂起，直到另一个goroutine执行Signal或者Broadcast，
	该goroutine才有机会重新唤醒，并尝试获取Locker，完成后续逻辑。
	 */
	c.Wait()
	m.Unlock()
	fmt.Println("Done")
}
