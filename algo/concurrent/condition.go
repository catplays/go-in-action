package concurrent

import (
	"fmt"
	"sync"
)

/**
- 条件变量是基于互斥锁的，它必须有互斥锁的支撑才能发挥作用。
- 条件变量的作用是，当共享资源发生变化的时候，可以通知被互斥锁阻塞的线程
典型例子：
阻塞队列
*/


var (
	mailBox uint8 = 0 // 1表示有信封 0表示没有
	lock sync.RWMutex
	sendCond = sync.NewCond(&lock)
	recvCond = sync.NewCond(lock.RLocker())
)

// 发邮件，一次只能一个人发
func SendMail() {
	lock.Lock()
	defer lock.Unlock()

	for mailBox == 1 {
		sendCond.Wait()
	}
	mailBox = 1
	fmt.Println("send mail.")
	recvCond.Signal()
}

func RecvMail() struct{}  {
	lock.RLock()
	defer lock.RUnlock()

	for mailBox == 0 {
		recvCond.Wait()
	}
	mailBox = 0
	fmt.Println("recv mail.")
	sendCond.Signal()
	return struct{}{}
}
