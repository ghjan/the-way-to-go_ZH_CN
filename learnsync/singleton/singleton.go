package singleton

import "sync"

/*
Go语言学习之sync包(临时对象池Pool、互斥锁Mutex、等待Cond)(the way to go)
https://blog.csdn.net/wangshubo1989/article/details/77966432

简单应用

通过sync实现单例
 */
type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
