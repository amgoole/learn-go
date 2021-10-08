package channel

import (
	"fmt"
	"time"
)

func ChanDemo() {
	//var c chan int // c == nil
	c := make(chan int)
	// 如果不存在下面协程，程序死锁 c为无缓冲chan 向c中写入数据时阻塞，main goroutine死锁
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	// 2并不是一定可以打印出来， 当协程取出2时 还未打印可能主协程已经结束，程序结束
}

func ChanDemo2() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go func(i int) {
			for {
				fmt.Printf("%d --- %c\n", i, <-channels[i])
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Second)
	go func() {}()

}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go func(i int) {
		for {
			fmt.Printf("%d --- %c\n", i, <-c)
		}
	}(i)
	return c
}

func ChanDemo3() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second * 2)
}

// chan 可以是有缓冲的
// 使用range获取chan的值 val := rang chan
// 发送方可以close chan 关闭后依然可以收到对应的0值
// val, ok := <-chan 关闭后 ok != nil
