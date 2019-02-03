package problem02

/*
题：
实现单例模式

什么是单例模式:
单例模式，也叫单件模式，是一种常见的软件设计模式。单例模式确保一个类只有一个实例，并提供

为什么要单例模式:
​有一些对象我们只需要一个，比方说：线程池(threadpool)、缓存(cache)、日志对象等等。事实上，这类对象只能有一个实例，如果制造出多个实例，就会导致许多问题产生，例如：程序的行为异常、资源使用过量，或者是不一致的结果。
​保证只有一个对象，省略了后续创建对象所花费的时间，也减轻了垃圾回收的压力(GC)

如何实现单例模式:
实现单例模式的关键是
1.保证对象只有一个
2.线程安全

sync.Once能确保实例化对象Do方法在多线程环境只运行一次,内部通过互斥锁实现
*/
import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
