package gorou

import (
	"fmt"
	"runtime"
	"time"
)

// Demo1 Demo 非抢占式 协程遇到IO交出执行权
func Demo1() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// Print 为io操作，会进行协程的切换
				fmt.Printf("goro %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Second)
}

func Demo2() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// 该条指令不会引起协程切换，仅有4个协程在运行（cpu=4）
				arr[i]++
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

func Demo3() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 变成闭包，当外部循环修改i=10时，内部数组出现越界
				arr[i]++
				// 使用 go run -race xx.go 检测数据冲突
			}
		}()
	}
	time.Sleep(time.Second)
}

func Demo4() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// 该条指令不会引起协程切换，仅有两个协程在运行（cpu=2）
				arr[i]++
				// 该方法会主动交出协程执行权限
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(arr) //同样产生数据冲突，协程在写arr print在打印
}
