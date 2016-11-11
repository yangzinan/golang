package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //去计算机cpu个数来聚顶一次启动线程的个数
	wg := sync.WaitGroup{}               //创建一个同步等待
	wg.Add(10)                           //同步等待10
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	var a int
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done() //执行函数结束后同步等待次数减一333333
}
